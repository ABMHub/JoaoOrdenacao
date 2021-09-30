package sort

import (
	"sortAlgorithms/util"
)

func Mergesort_F(arr []util.T, begin, end int, cmp func(util.T, util.T) bool) {

	if begin >= end {
		return
	}
	var mid int = (begin + end) / 2

	gr := func() {
		Mergesort_F(arr, begin, mid, cmp)
	}
	gr2 := func() {
		Mergesort_F(arr, mid+1, end, cmp)
	}
		
	util.Semaforo(1,gr, gr2)

	merge(arr, begin, mid, end, cmp)
}