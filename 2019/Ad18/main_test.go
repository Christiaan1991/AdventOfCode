package main

import "testing"

func TestIsKey(t *testing.T) {
	iskey := isKey("a")

	if !iskey {
		t.Errorf("is not a key!")
	}
}

