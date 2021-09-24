package main

import (
	"fmt"
	//	"os"
	"sortAlgorithms/sort"
	"sortAlgorithms/util"
	//"time"
)

func main() {
	arr := []util.T{5, 2, 1, -12312, 2312312}

	util.SetThreadLimit(6)

	sort.Quicksort_F(arr, 0, len(arr)-1, func(a, b util.T) bool {
		return a.(int) < b.(int)
	})
	for i := range arr {
		fmt.Println(arr[i])
	}
}
