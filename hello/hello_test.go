package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {

		// t.Helper() is needed to tell the test suite that this method is a helper.
		// By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to the Worlds Greatest Detective", func(t *testing.T) {
		got := Hello("Bruce", "")
		want := "Hello, Bruce"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying 'Hello, Sir' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, Sir"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("ウェインさん", "japanese")
		want := "こんにちは、ウェインさん"

		assertCorrectMessage(t, got, want)
	})
}
