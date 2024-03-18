package integers

import (
	"fmt"
	"testing"
)

// Add takes two integers and returns the sum of these
func TestAdder(t *testing.T) {
	want := 5
	got := Add(2, 3)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
