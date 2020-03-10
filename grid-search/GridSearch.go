package gridsearch

import "fmt"

// GridSearch return YES if found pattern P inside
// grid G
func GridSearch(G []string, P []string) string {
	R := len(G)
	C := len(G[0])
	r := len(P)
	c := len(P[0])
	fmt.Println(R, C, r, c)

	return "NO"
}
