package main

import (
	"strings"
	"testing"
	"time"
)

func TestRenderTerminalNodesSortsByName(t *testing.T) {
	now := time.Now().UnixNano()
	out := renderTerminalNodes([]NodeStats{
		{Name: "z-node", Version: appVersion, CPU: []float64{1}, MemTotal: 1, UpdatedAt: now},
		{Name: "a-node", Version: appVersion, CPU: []float64{1}, MemTotal: 1, UpdatedAt: now},
	})

	first := strings.Index(out, "a-node")
	second := strings.Index(out, "z-node")
	if first < 0 || second < 0 {
		t.Fatalf("missing rendered nodes:\n%s", out)
	}
	if first > second {
		t.Fatalf("nodes were not sorted by name:\n%s", out)
	}
}
