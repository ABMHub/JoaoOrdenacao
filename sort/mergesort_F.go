package sort

import (
	"sortAlgorithms/util"
	"runtime"
	"sync"
)

func Mergesort_F(arr []util.T, begin, end int, cmp func(util.T, util.T)bool) {
	if begin == 0 && end == len(arr)-1{
		runtime.GOMAXPROCS(3) // aloca 
	}

	if begin >= end {
		return
	}

	var tg sync.WaitGroup
	var mid int = (begin + end) / 2

	tg.Add(1)
	
	go func(arr []util.T, begin,mid, end int, cmp func(util.T, util.T)bool){
		defer tg.Done()
		Mergesort_F(arr, begin, mid,cmp)
		Mergesort_F(arr, mid+1, end,cmp)
	}(arr,begin,mid,end,cmp)

	tg.Wait()

	merge_F(arr, begin, mid, end,cmp)
}

func merge_F(arr []util.T, begin, mid, end int,cmp func(util.T, util.T)bool) {
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