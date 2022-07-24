package integers

import "testing"

func TestAddr(t *testing.T) {
	sum := Add(3, 2)
	expected := 5

	if sum != expected {
		t.Errorf("expected %q but got %q", expected, sum)
	}
}
