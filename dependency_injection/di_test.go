package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Bruce Wayne")

	got := buffer.String()
	want := "Hello, Bruce Wayne"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
