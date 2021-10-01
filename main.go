package main

import (
	"fmt"
	"log"
	"os"
	"sortAlgorithms/sort"
	"sortAlgorithms/util"
)

// 	"os"
// "sortAlgorithms/sort"
// "sortAlgorithms/util"
//"time"

func main() {
	// arr := []util.T{5, 2, 1, -12312, 2312312}

	// util.SetThreadLimit(6)

	// sort.Quicksort_F(arr, 0, len(arr)-1, func(a, b util.T) bool {
	// 	return a.(int) < b.(int)
	// })
	// for i := range arr {
	// 	fmt.Println(arr[i])
	// }
	size := 4
	sort.Merge_Files(util.ReadIntegers, "merge-sort", size, 1000, util.CompareInt)
	// file, err := os.Open("integerscpp.bin")

	file, err := os.Open("temp/out1.bin") // abre arquivo
	if err != nil {                       // se der erro cancela tudo
		log.Fatal("Erro na leitura do arquivo binario com os inteiros a serem ordenados", err) //
		defer file.Close()                                                                     //
	}

	fmt.Println(util.ReadIntegers(file, 10))
}
