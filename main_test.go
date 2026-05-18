package main

import (
	"os"
	"testing"
)

func withoutNodeNameEnv(t *testing.T) {
	t.Helper()
	old, hadOld := os.LookupEnv(envNodeName)
	if err := os.Unsetenv(envNodeName); err != nil {
		t.Fatalf("unset %s: %v", envNodeName, err)
	}
	t.Cleanup(func() {
		if hadOld {
			_ = os.Setenv(envNodeName, old)
		} else {
			_ = os.Unsetenv(envNodeName)
		}
	})
}

func TestNodeNameFromEnvUsesHostnameByDefault(t *testing.T) {
	withoutNodeNameEnv(t)

	got, err := nodeNameFromEnv("host-a")
	if err != nil {
		t.Fatalf("nodeNameFromEnv: %v", err)
	}
	if got != "host-a" {
		t.Fatalf("name = %q, want host-a", got)
	}
}

func TestNodeNameFromEnvUsesOverride(t *testing.T) {
	t.Setenv(envNodeName, "test-node")

	got, err := nodeNameFromEnv("host-a")
	if err != nil {
		t.Fatalf("nodeNameFromEnv: %v", err)
	}
	if got != "test-node" {
		t.Fatalf("name = %q, want test-node", got)
	}
}

func TestNodeNameFromEnvRejectsWhitespaceOverride(t *testing.T) {
	for _, value := range []string{" test-node", "test-node ", "test node", "\t"} {
		t.Run(value, func(t *testing.T) {
			t.Setenv(envNodeName, value)
			if got, err := nodeNameFromEnv("host-a"); err == nil {
				t.Fatalf("name = %q, want error", got)
			}
		})
	}
}
