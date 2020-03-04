package biggerisgreater

import (
	"fmt"
	"strings"
)

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "z"}

// Complete the biggerIsGreater function below.
func BiggerIsGreater(w string) string {
	w = " sod "
	z := stringToArray(w)
	fmt.Println(findIndexInArr(alphabet, z[1]))
	return w
}

func findIndexInArr(arr []string, el string) int {
	for i, v := range arr {
		if v == el {
			return i
		}
	}
	return -1
}

func stringToArray(arr string) []string {
	return strings.Split(strings.TrimSpace(arr), "")
}
