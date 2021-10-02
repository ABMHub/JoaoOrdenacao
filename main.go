package main

import (
	// "fmt"
	// "log"
	// "os"

	"fmt"

	"sortAlgorithms/sort"
	"sortAlgorithms/util"
)

func main() {
	size := 4
	sort.Merge_Files(util.ReadIntegers, "merge-sort", size, util.CompareInt)
	util.SetThreadLimit(4)
	//util.GenerateFiles(250000000)

	// file1, err1 := os.Open("IntegersGo.bin") // abre arquivo
	// if err1 != nil {                         // se der erro cancela tudo
	// 	log.Fatal("Erro na leitura do arquivo binario com os inteiros a serem ordenados", err1) //
	// 	defer file1.Close()                                                                     //
	// }

	//fmt.Println(util.ReadIntegers(file1, 10))
	fmt.Println("oi")
}
