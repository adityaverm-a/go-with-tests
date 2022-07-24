package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Mr. Wayne")
	want := "Hello, Mr. Wayne"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
