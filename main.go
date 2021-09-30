package main

import (
	"fmt"
	"os"
	"sortAlgorithms/sort"
	"sortAlgorithms/util"
	"time"
	"encoding/binary"
)


func main() {
	file,_ := os.Open("ints.txt")
	util.SetThreadLimit(8)
	temp := make([]byte, 4)
	arr := make([]util.T,1e6)
	for i := 0; i < 1e6; i++{
		file.Read(temp)
		arr[i] = int(binary.BigEndian.Uint32(temp))
	}
	fmt.Println(arr[len(arr)-1])
	start := time.Now()
	for i := 0; i < 10; i++{
		sort.Mergesort_P(arr,func(a, b util.T) bool {
			return a.(int) < b.(int)
		})
		fmt.Println(i)
	}
	fmt.Println(time.Since(start))
}

