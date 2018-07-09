package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello test circle"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
