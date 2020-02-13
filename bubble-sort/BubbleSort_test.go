package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	slices := []struct {
		a []int
		s []int
		e []int
	}{
		{[]int{1, 2}, []int{}, []int{1, 2}},
		{[]int{}, []int{}, []int{}},
		{[]int{6, 3, 0, 1, 9}, []int{}, []int{0, 1, 3, 6, 9}},
		{[]int{6, 5, 3, 1, 8, 7, 2, 4}, []int{}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}
	for _, slice := range slices {
		total := BubbleSort(slice.a, slice.s)
		if !reflect.DeepEqual(total, slice.e) {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", slice.a, slice.s, total, slice.e)
		}
	}
}
