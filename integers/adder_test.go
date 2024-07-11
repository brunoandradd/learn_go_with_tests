package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	want := 4
     
	if sum != want {
		t.Errorf("expected: %d, got: %d", want, sum)
	}
}

func ExampleAdd() {
	sum := Add(5, 9)
	fmt.Println(sum) // output: 14
}
