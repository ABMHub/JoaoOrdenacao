package bsort

import (
	"fmt"
	"sortAlgorithms/util"
)

func Bubblesort(v []util.T, cmp func(util.T, util.T) bool) {
	var flag bool = false
	var v_size int = len(v)

	for !flag {
		flag = true
		for i := 0; i < v_size-1; i++ {
			if cmp(v[i], v[i+1]) {
				flag = false
				v[i], v[i+1] = v[i+1], v[i]
			}
		}
	}

	for i := 0; i < v_size; i++ {
		fmt.Printf("%s ", v[i])
	}
}
