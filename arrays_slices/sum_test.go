package arraysslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		want := 15
		got := Sum(numbers)
		if want != got {
			t.Errorf("expected %v, found %v", want, got)
		}
	})

}
