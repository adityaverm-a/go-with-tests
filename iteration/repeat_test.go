package iteration

import (
	"fmt"
	"testing"
)

// Run, go test -v
func ExampleRepeat() {
	repeated := Repeat("c", 3)
	fmt.Println(repeated)
	// Output: ccc
}

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, repeated, expected string) {
		t.Helper()

		if expected != repeated {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("Repeat any given characters 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Repeat any given characters as many times as provided", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		assertCorrectMessage(t, repeated, expected)
	})
}

// To run the benchmarks do, go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", 6)
	}
}
