package main

import "testing"

func TestMainFunc(t *testing.T) {
	s := hello()

	if s != "Hello World" {
		t.Errorf("Eroor ")
	}
}
