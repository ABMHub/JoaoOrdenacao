package main

import (
	"fmt"
	"os"
	"sortAlgorithms/sort"
	"sortAlgorithms/util"
	"encoding/binary"
)


func main() {
	file,_ := os.Open("ints.txt")
	util.SetThreadLimit(6)
	temp := make([]byte, 4)
	arr := make([]util.T,1e6)
	for i := 0; i < 1e6; i++{
		file.Read(temp)
		arr[i] = int(binary.BigEndian.Uint32(temp))
	}
	sort.Mergesort_P(arr, 0, len(arr)-1, func(a, b util.T)bool{
		return a.(int) > b.(int)
	})
	fmt.Println(arr)
}

