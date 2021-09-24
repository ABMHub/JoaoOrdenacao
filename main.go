package main

import (
	"fmt"
	"os"
	"sortAlgorithms/sort"
	"sortAlgorithms/util"
	"time"
)

func main() {
	var v [1000000]util.T
	//var v_size int = len(v)
	file, err := os.Open("ints.txt")
	if err != nil{
		fmt.Printf("DEU ERRADO KK");
		return
	}
	var temp int
	for i := 0; i < 1e6; i++{
		fmt.Fscanf(file,"%d\n", &temp)
		v[i]=temp
	}
	v2 := v[:]
	start := time.Now()
	for i := 1; i <= 10; i++{
		sort.Mergesort(v2,0, len(v)-1, util.CompareInt)
		fmt.Printf("%d\n", i)
	}
	end := time.Now()
	fmt.Printf("tempo : %d",end.Sub(start));
}
