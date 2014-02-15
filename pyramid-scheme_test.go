package main

import "testing"

func TestAdd(t *testing.T) {
	const n, m = 3, 3
	const out = 6
	if x := Add(n, m); x != out {
		t.Errorf("Add(%v, %v) = %v, want %v", n, m, x, out)
	}
}
