package gridsearch

import (
	"strings"
)

// GridSearch return YES if found pattern P inside
// grid G
func GridSearch(G []string, P []string) string {
	C := len(G)
	c := len(P)
Loop:
	for i, v := range G {
		if i+c > C {
			break Loop
		}
		for _, j := range findAllIndex(v, P[0]) {
		SecondLoop:
			for l, sv := range P {
				comp := G[i+l][j:]
				if t := strings.Index(comp, sv); t != 0 {
					break SecondLoop
				} else if l == len(P)-1 {
					return "YES"
				}
			}
		}
	}
	return "NO"
}

func findAllIndex(str, substr string) []int {
	var fI int
	s := len(str)
	o := []int{}
Loop:
	for i := 0; i < s; i++ {
		fI = strings.Index(str, substr)
		if fI == -1 {
			break Loop
		}
		o = append(o, fI+i)
		str = removeElementByIndex(str, fI)
	}
	return o
}

func removeElementByIndex(v string, i int) string {
	s := []rune(v)
	return string(append(s[:i], s[i+1:]...))
}
