package gridsearch

import (
	"reflect"
	"testing"
)

func TestGridSearch(t *testing.T) {
	patterns := []struct {
		g []string
		p []string
		e string
	}{
		{
			[]string{"7283455864", "6731158619", "8988242643", "3830589324", "2229505813", "5633845374", "6473530293", "7053106601", "0834282956", "4607924137"},
			[]string{"9505", "3845", "3530"},
			"YES",
		},
		{
			[]string{"400453592126560", "114213133098692", "474386082879648", "522356951189169", "887109450487496", "252802633388782", "502771484966748", "075975207693780", "511799789562806", "404007454272504", "549043809916080", "962410809534811", "445893523733475", "768705303214174", "650629270887160"},
			[]string{"99", "99"},
			"NO",
		},
		{
			[]string{"1", "1"},
			[]string{"1", "1"},
			"YES",
		},
	}
	for _, str := range patterns {
		total := GridSearch(str.g, str.p)
		if !reflect.DeepEqual(total, str.e) {
			t.Errorf("Operation over (%v) with patter %v failed. got: %v, want: %v.", str.g, str.p, total, str.e)
		}
	}
}
