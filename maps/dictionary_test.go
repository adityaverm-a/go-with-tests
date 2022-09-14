package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is just a test!"}

	got := Dictionary.Search(dictionary, "test")
	want := "This is just a test!"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
