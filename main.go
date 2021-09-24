package main

import (
	"fmt"
	"sortAlgorithms/sort"
	"sortAlgorithms/util"
)

func main() {
	v := []util.T{"ab", "csc", "asrasr", "kitopott"}
	var v_size int = len(v)

	fmt.Println()

	sort.Mergesort_F(v[:],0, len(v)-1, util.CompareString)
	for i := 0; i < v_size; i++ {
		fmt.Printf("%s ", v[i])
	}

}
