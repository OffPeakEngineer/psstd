package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/cockroachdb/pebble/v2"
)

func openTestDB(t *testing.T) *pebble.DB {
	t.Helper()
	db, err := pebble.Open(t.TempDir(), &pebble.Options{})
	if err != nil {
		t.Fatalf("open pebble: %v", err)
	}
	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Fatalf("close pebble: %v", err)
		}
	})
	return db
}

func mustGetNode(t *testing.T, db *pebble.DB, name string) NodeStats {
	t.Helper()
	b, closer, err := db.Get(keyFor(name))
	if err != nil {
		t.Fatalf("get node %s: %v", name, err)
	}
	defer closer.Close()
	var s NodeStats
	if err := json.Unmarshal(b, &s); err != nil {
		t.Fatalf("unmarshal node %s: %v", name, err)
	}
	return s
}

func TestMergeAcceptsOnlineDifferentVersion(t *testing.T) {
	db := openTestDB(t)
	incoming := NodeStats{Name: "old-node", Version: "v1.0.0", UpdatedAt: time.Now().UnixNano()}

	if err := dbMergeLWW(db, incoming, "v2.0.0"); err != nil {
		t.Fatalf("merge online different version: %v", err)
	}
	got := mustGetNode(t, db, "old-node")
	if got.Version != incoming.Version {
		t.Fatalf("version = %q, want %q", got.Version, incoming.Version)
	}
}

func TestMergeRejectsOfflineDifferentVersion(t *testing.T) {
	db := openTestDB(t)
	incoming := NodeStats{Name: "old-node", Version: "v1.0.0", UpdatedAt: 0}

	if err := dbMergeLWW(db, incoming, "v2.0.0"); err != errStaleVersion {
		t.Fatalf("merge error = %v, want %v", err, errStaleVersion)
	}
	if _, closer, err := db.Get(keyFor("old-node")); err == nil {
		closer.Close()
		t.Fatal("stale offline node was stored")
	}
}

func TestPurgeOfflineDifferentVersionKeepsOnlineDifferentVersion(t *testing.T) {
	db := openTestDB(t)
	if err := dbSet(db, NodeStats{Name: "offline-old", Version: "v1.0.0", UpdatedAt: 0}); err != nil {
		t.Fatalf("set offline-old: %v", err)
	}
	if err := dbSet(db, NodeStats{Name: "online-old", Version: "v1.0.0", UpdatedAt: time.Now().UnixNano()}); err != nil {
		t.Fatalf("set online-old: %v", err)
	}

	if err := purgeOfflineDifferentVersion(db, "v2.0.0"); err != nil {
		t.Fatalf("purge: %v", err)
	}
	if _, closer, err := db.Get(keyFor("offline-old")); err == nil {
		closer.Close()
		t.Fatal("offline-old was not purged")
	}
	got := mustGetNode(t, db, "online-old")
	if got.Version != "v1.0.0" {
		t.Fatalf("online-old version = %q, want v1.0.0", got.Version)
	}
}
