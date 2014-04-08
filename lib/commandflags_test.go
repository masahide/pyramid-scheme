package lib

import (
	"testing"
)


func TestShowVersion(t *testing.T) {
	if ShowVersion() != "pyramid-scheme version: -" {
		t.Errorf("showVersion() = [%s], want %v", ShowVersion(), "pyramid-scheme version: -")
	}
}

func TestUsage(t *testing.T) {
	Usage()
}
