package main

import (
	"fmt"
	"log"
	"os"

	"sortAlgorithms/sort"
	"sortAlgorithms/util"
	// "time"
	// "github.com/cheggaaa/pb"
)

func main() {
	size := 4
	number_of_processors := 12
	sort.Merge_Files("Integerscpp2.bin", "quick-sort", size, 12, number_of_processors, util.ReadIntegers, util.CompareInt, util.FragmentBin, util.WriteIntegers)
	util.SetThreadLimit(1)

	file1, err1 := os.Open("Sorted.bin") // abre arquivo
	if err1 != nil {                     // se der erro cancela tudo
		log.Fatal("Erro na leitura do arquivo binario com os inteiros a serem ordenados", err1) //
		defer file1.Close()                                                                     //
	}

	fmt.Println(util.ReadIntegers(file1, 100))
	fmt.Println("oi")
}
