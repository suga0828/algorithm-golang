package biggerisgreater

import (
	"reflect"
	"testing"
)

func TestBiggerIsGreater(t *testing.T) {
	strings := []struct {
		i string
		e string
	}{
		{"ab", "ba"},
		{"bb", "no answer"},
		{"hefg", "hegf"},
		{"dhck", "dhkc"},
		{"lmno", "lmon"},
		{"dcba", "no answer"},
		{"dcbb", "no answer"},
		{"abdc", "acbd"},
		{"abcd", "abdc"},
		{"fedcbabcd", "fedcbabdc"},
	}
	for _, str := range strings {
		total := BiggerIsGreater(str.i)
		if !reflect.DeepEqual(total, str.e) {
			t.Errorf("Operation over (%v) failed. got: %v, want: %v.", str.i, total, str.e)
		}
	}
}
