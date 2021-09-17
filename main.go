package main

import (
	"fmt"
	"sortAlgorithms/sort"
	"sortAlgorithms/util"
)

func main() {
	v := []util.T{"ab", "csc", "asrasr", "kitopott"}
	var v_size int = len(v)

	for i := 0; i < v_size; i++ {
		fmt.Printf("%s ", v[i])
	}
	fmt.Println()

	sort.Bubblesort(v[:], util.CompareString)
}
