package sort

import (
	"sortAlgorithms/util"
)

func Mergesort(arr []util.T, begin, end int, cmp func(util.T, util.T) bool) {
	if begin >= end {
		return
	}

	var mid int = (begin + end) / 2
	Mergesort(arr, begin, mid, cmp)
	Mergesort(arr, mid+1, end, cmp)
	merge(arr, begin, mid, end, cmp)
}

func merge(arr []util.T, begin, mid, end int, cmp func(util.T, util.T) bool) {
	temp := make([]util.T, end-begin+1)

	var i int = begin
	var j int = mid + 1
	var k int = 0

	for i <= mid && j <= end {
		if cmp(arr[i], arr[j]) {
			temp[k] = arr[i]
			k++
			i++
		} else {
			temp[k] = arr[j]
			k++
			j++
		}
	}

	for i <= mid {
		temp[k] = arr[i]
		k++
		i++
	}

	for j <= end {
		temp[k] = arr[j]
		k++
		j++
	}

	for l := begin; l <= end; l++ {
		arr[l] = temp[l-begin]
	}
}
