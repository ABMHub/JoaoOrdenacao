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

	file, err := os.Open("temp/out0.bin") // abre arquivo
	if err != nil {                       // se der erro cancela tudo
		log.Fatal("Erro na leitura do arquivo binario com os inteiros a serem ordenados", err) //
		defer file.Close()                                                                     //
	}

	// fmt.Println("++++++++++++++++++++++++++++++++", util.ReadIntegers(file, 1000))

	file1, err1 := os.Open("temp/out1.bin") // abre arquivo
	if err1 != nil {                        // se der erro cancela tudo
		log.Fatal("Erro na leitura do arquivo binario com os inteiros a serem ordenados", err1) //
		defer file1.Close()                                                                     //
	}

	// fmt.Println("\n\n\n\n<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<", util.ReadIntegers(file1, 1000))

	sort.Merge_arrays(util.ReadIntegers, util.CompareInt, file1, file, 1000)

	file2, err2 := os.Open("output.bin") // abre arquivo
	if err2 != nil {                     // se der erro cancela tudo
		log.Fatal("Erro na leitura do arquivo binario com os inteiros a serem ordenados", err2) //
		defer file2.Close()                                                                     //
	}

	fmt.Println("\n\n\n\n>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", util.ReadIntegers(file2, 2000))
}
