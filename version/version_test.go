package version

import (
	"testing"
)

func TestVersion(t *testing.T) {

	GITCOMMIT = "hoge"
	if GITCOMMIT != "hoge" {
		t.Errorf("%+v, want %v", GITCOMMIT, "hoge")
	}
	VERSION = "hoge"
	if VERSION != "hoge" {
		t.Errorf("%+v, want %v", VERSION, "hoge")
	}
}
