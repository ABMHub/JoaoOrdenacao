package sort

import (
	"sortAlgorithms/util"
	"sync"
)
var tg sync.WaitGroup

func Quicksort_H(arr []util.T, low, high int, cmp func(util.T, util.T) bool, p_threads int){
	if low >= high {
		return
	}
	tg.Add(1)
	var pivot_index int = partition(arr, low, high, cmp)
	if p_threads == 0{
		Quicksort(arr, low, pivot_index-1, cmp)
		Quicksort(arr, pivot_index+1, high, cmp)
	}else if p_threads == 1{
		go Quicksort(arr, low, pivot_index-1, cmp)
		Quicksort(arr, pivot_index+1, high, cmp)
	}else{
		p_threads -= 1;
		go Quicksort_H(arr, low, pivot_index-1, cmp,p_threads/2 + (p_threads%2))
	    Quicksort_H(arr, pivot_index+1, high, cmp,p_threads/2)
	}
	tg.Done()
}

func Quicksort_P(arr []util.T, low, high int, cmp func(util.T, util.T) bool) {
	Quicksort_H(arr,low,high,cmp,util.GetThreadLimit())
	tg.Wait()
}