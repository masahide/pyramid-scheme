package main

import (
	"testing"
)

func TestShowVersion(t *testing.T) {
	if showVersion() != "pyramid-scheme version: -" {
		t.Errorf("showVersion() = [%s], want %v", showVersion(), "pyramid-scheme version: -")
	}
}
