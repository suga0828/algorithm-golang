package biggerisgreater

import (
	"sort"
	"strings"
)

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

// BiggerIsGreater return the smallest lexicographically
// higher string possible from the given string or "no answer".
func BiggerIsGreater(w string) string {
	z := stringToAlphabeticIndex(w)

	for i := len(z) - 1; i > 0; i-- {
		if z[i-1] < z[i] {
			b, bI := smallestGreaterThanAndIndex(z[i-1:len(z)], z[i-1], i-1)
			z[bI] = z[i-1]
			z[i-1] = b
			pp := z[i:len(z)]
			sort.Ints(pp)
			pz := append(z[0:i], pp...)
			nz := alphabeticIndexToString(pz)
			return nz
		}
	}
	return "no answer"
}

func mapStrInt(vs []string, f func(vs []string, t string) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(alphabet, v)
	}
	return vsm
}

func mapIntStr(vs []int) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = alphabet[v]
	}
	return vsm
}

func indexStr(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func elFromIndex(vs []string, t int) int {
	for i := range vs {
		if i == t {
			return i
		}
	}
	return -1
}

func indexInt(vs []int, t int) int {
	for i := len(vs) - 1; i > 0; i-- {
		if vs[i] == t {
			return i
		}
	}
	return -1
}

func stringToAlphabeticIndex(str string) []int {
	arr := strings.Split(strings.TrimSpace(str), "")
	return mapStrInt(arr, indexStr)
}

func alphabeticIndexToString(arr []int) string {
	narr := mapIntStr(arr)
	return strings.Join(narr, "")
}

func smallestGreaterThanAndIndex(arr []int, n int, i int) (int, int) {
	var s int
	var l int
	for k, e := range arr {
		if e > n {
			if s == 0 && e != n {
				s = e
				l = k + i
			}
			if e < s {
				s = e
				l = k + i

			}
		}
	}
	return s, l
}
