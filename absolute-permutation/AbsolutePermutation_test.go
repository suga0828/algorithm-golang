package absolutepermutation

import (
	"reflect"
	"testing"
)

func TestAbsolutePermutation(t *testing.T) {
	integers := []struct {
		n int32
		k int32
		e []int32
	}{
		{2, 1, []int32{2, 1}},
		{3, 0, []int32{1, 2, 3}},
		{3, 2, []int32{-1}},
		{4, 2, []int32{3, 4, 1, 2}},
		{6, 3, []int32{4, 5, 6, 1, 2, 3}},
		{6, 2, []int32{-1}},
		{8, 2, []int32{3, 4, 1, 2, 7, 8, 5, 6}},
	}
	for _, integer := range integers {
		total := AbsolutePermutation(integer.n, integer.k)
		if !reflect.DeepEqual(total, integer.e) {
			t.Errorf("Operation over %v %v failed. got: %v, \n want: %v.", integer.n, integer.k, total, integer.e)
		}
	}
}
