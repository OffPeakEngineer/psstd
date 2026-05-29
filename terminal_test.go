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

func TestRenderTerminalGridPacksStableRows(t *testing.T) {
	now := time.Now().UnixNano()
	out := renderTerminalGrid([]NodeStats{
		{Name: "c-node", Version: appVersion, CPU: []float64{1}, MemTotal: 1, UpdatedAt: now},
		{Name: "a-node", Version: appVersion, CPU: []float64{1}, MemTotal: 1, UpdatedAt: now},
		{Name: "b-node", Version: appVersion, CPU: []float64{1}, MemTotal: 1, UpdatedAt: now},
	}, terminalCellWidth*2+terminalCellGap)

	first := strings.Index(out, "a-node")
	second := strings.Index(out, "b-node")
	third := strings.Index(out, "c-node")
	if first < 0 || second < 0 || third < 0 {
		t.Fatalf("missing rendered nodes:\n%s", out)
	}
	if first > second || second > third {
		t.Fatalf("grid nodes were not sorted by name:\n%s", out)
	}
	if strings.Count(out, "╭") != 3 {
		t.Fatalf("grid did not render three bordered cells:\n%s", out)
	}
}

func TestTerminalColumns(t *testing.T) {
	if got := terminalColumns(terminalCellWidth - 1); got != 1 {
		t.Fatalf("narrow columns = %d, want 1", got)
	}
	if got := terminalColumns(terminalCellWidth*2 + terminalCellGap); got != 2 {
		t.Fatalf("wide columns = %d, want 2", got)
	}
}

func TestTerminalHeaderExplainsViewerModeAndControls(t *testing.T) {
	now := time.Date(2026, 5, 29, 12, 34, 56, 0, time.UTC)

	plain := terminalHeader(2, now, false)
	for _, want := range []string{"viewer-only", "2 node(s)", "2026-05-29T12:34:56Z", "refreshes every 2s"} {
		if !strings.Contains(plain, want) {
			t.Fatalf("plain header missing %q: %s", want, plain)
		}
	}

	interactive := terminalHeader(2, now, true)
	for _, want := range []string{"viewer-only", "q quit", "r refresh", "Ctrl-C quit"} {
		if !strings.Contains(interactive, want) {
			t.Fatalf("interactive header missing %q: %s", want, interactive)
		}
	}
}

func TestTerminalKeyAction(t *testing.T) {
	cases := []struct {
		key  byte
		want terminalAction
	}{
		{key: 'q', want: terminalActionQuit},
		{key: 'Q', want: terminalActionQuit},
		{key: 3, want: terminalActionQuit},
		{key: 'r', want: terminalActionRefresh},
		{key: 'R', want: terminalActionRefresh},
		{key: 'x', want: terminalActionNone},
	}
	for _, tc := range cases {
		if got := terminalKeyAction(tc.key); got != tc.want {
			t.Fatalf("key %q action = %d, want %d", tc.key, got, tc.want)
		}
	}
}
