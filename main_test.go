package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(1, 5)
	if total != 6 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 6)
	}
}
