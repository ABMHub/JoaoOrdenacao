package main

import (
	"fmt"
	"log"
	"os"

	"sortAlgorithms/sort"
	"sortAlgorithms/util"
)

func main() {
	// External
	elem_size := 4
	number_of_processors := 12
	batch_size := 10
	util.SetThreadLimit(1)

	// Chama a ordenacao
	sort.Merge_Files("Integerscpp2.bin", "quick-sort", elem_size, batch_size, number_of_processors, util.ReadIntegers, util.CompareInt, util.FragmentBin, util.WriteIntegers)

	// Primeiros 100 elementos do arquivo a ser ordenado
	fmt.Print("\n\n100 primeiros inteiros a serem ordenados:\n\n")
	file1, err1 := os.Open("Integerscpp2.bin") // abre arquivo
	if err1 != nil {                           // se der erro cancela tudo
		log.Fatal("Erro na leitura do arquivo binario com os inteiros a serem ordenados", err1) //
		defer file1.Close()                                                                     //
	}
	fmt.Println(util.ReadIntegers(file1, 100))

	// Mostra os 100 primeiros elementos depois de o arquivo ter sido ordenado
	fmt.Print("\n\n100 primeiros inteiros depois de ordenados:\n\n")
	file1, err1 = os.Open("Sorted.bin") // abre arquivo
	if err1 != nil {                    // se der erro cancela tudo
		log.Fatal("Erro na leitura do arquivo binario com os inteiros ordenados", err1) //
		defer file1.Close()                                                             //
	}
	fmt.Println(util.ReadIntegers(file1, 100))

	/* ###################################################################################### */

	// Json

	// file1, err1 := os.Open("dump.json")
	// if err1 != nil {
	// 	log.Fatal("Erro na leitura do arquivo binario com os inteiros a serem ordenados", err1)
	//  	defer file1.Close()
	// }
	// sort.SortJson(file1, util.CompareLikes)
}
