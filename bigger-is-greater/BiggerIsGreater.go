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
			b := smallestGreaterThan(z[i-1:len(z)], z[i-1])
			bI := indexInt(z, b)
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
	for i, v := range vs {
		if v == t {
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

func smallestGreaterThan(arr []int, n int) int {
	gArr := []int{}
	for _, e := range arr {
		if e > n {
			gArr = append(gArr, e)
		}
	}
	s := gArr[0]
	for _, v := range gArr {
		if v < s {
			s = v
		}
	}
	return s
}
