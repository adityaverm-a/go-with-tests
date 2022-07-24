package integers

import (
	"fmt"
	"testing"
)

// Run, go test -v
func ExampleAdd() {
	sum := Add(6, 5)
	fmt.Println(sum)
	// Output: 11
}

func TestAddr(t *testing.T) {
	sum := Add(3, 2)
	expected := 5

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}
