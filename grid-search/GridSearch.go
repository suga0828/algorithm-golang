package gridsearch

import (
	"strings"
)

// GridSearch return YES if found pattern P inside
// grid G
func GridSearch(G []string, P []string) string {
	C := len(G[0])
	c := len(P[0])
	for i, v := range G {
		if p := strings.Index(v, P[0]); p != -1 {
			if i+c <= C {
				k := strings.Index(v, P[0])
			Loop:
				for l, sv := range P {
					if strings.Index(string(G[i+l]), sv) != k {
						break Loop
					} else if l == len(P)-1 {
						return "YES"
					}
				}
			}
		}
	}
	return "NO"
}
