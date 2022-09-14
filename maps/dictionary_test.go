package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is just a test!"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := Dictionary.Search(dictionary, "test")
		want := "This is just a test!"

		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
