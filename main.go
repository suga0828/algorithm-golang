package main

import (
	"fmt"

	bubblesort "github.com/suga0828/go-pkgs/bubble-sort"

	biggerisgreater "github.com/suga0828/go-pkgs/bigger-is-greater"
)

func main() {
	fmt.Println(biggerisgreater.BiggerIsGreater("sod"))
	fmt.Println(bubblesort.BubbleSort([]int{2, 1}, []int{}))
}
