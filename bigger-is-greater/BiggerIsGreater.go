package biggerisgreater

import (
	"sort"
)

type runeSlice []rune

func (p runeSlice) Len() int           { return len(p) }
func (p runeSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p runeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// BiggerIsGreater return the smallest lexicographically
// higher string possible from the given string or "no answer".
func BiggerIsGreater(w string) string {
	z := []rune(w)

	for i := len(z) - 1; i > 0; i-- {
		if int(z[i-1]) < int(z[i]) {
			b, bI := smallestGreaterThanAndIndex(z[i-1:len(w)], z[i-1], i-1)
			z[bI] = z[i-1]
			z[i-1] = b
			var pp runeSlice = z[i:len(z)]
			sort.Sort(pp)
			pz := append(z[0:i], pp...)
			return string(pz)
		}
	}
	return "no answer"
}

func smallestGreaterThanAndIndex(arr []rune, n rune, i int) (rune, int) {
	var s rune
	var l int
	for k, e := range arr {
		if int(e) > int(n) {
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
