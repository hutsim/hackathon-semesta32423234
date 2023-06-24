package main

import "testing"

func TestRun(t *testing.T) {
	expected := "Hackathon SEMESTA - System Administrator"
	got := run()
	if expected != got {
		t.Fatalf("expected %v got %v", expected, got)
	}
}
