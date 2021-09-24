package sort

import (
	"sortAlgorithms/util"
)

func partition(arr []util.T, low, high int, cmp func(util.T, util.T) bool) int {
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

func Quicksort(arr []util.T, low, high int, cmp func(util.T, util.T) bool) {
	if low >= high {
		return
	}

	var pivot_index int = partition(arr, low, high, cmp)
	Quicksort(arr, low, pivot_index-1, cmp)
	Quicksort(arr, pivot_index+1, high, cmp)
}
