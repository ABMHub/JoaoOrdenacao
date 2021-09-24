package sort

import (
	"sortAlgorithms/util"
)

func partition_F(arr []util.T, low, high int, cmp func(util.T, util.T) bool) int {
	var index int = low
	var pivot int = high

	for i := low; i < high; i++ {
		if cmp(arr[i], arr[pivot]) {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}

	arr[pivot], arr[index] = arr[index], arr[pivot]

	return index
}

func Quicksort_F(arr []util.T, low, high int, cmp func(util.T, util.T) bool) {
	if low >= high {
		return
	}

	var pivot_index int = partition_F(arr, low, high, cmp)

	gr := func() {
		Quicksort_F(arr, low, pivot_index-1, cmp)
		Quicksort_F(arr, pivot_index+1, high, cmp)
	}
	util.Semaforo(gr)
}
