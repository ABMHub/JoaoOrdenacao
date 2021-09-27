package sort

import (
	"fmt"
	"log"
	"math"
	"os"
	"sortAlgorithms/util"
	"sync"
)

var queueLock = &sync.Mutex{}

func Merge_arrays(cmp func(util.T, util.T) bool, file1, file2 *os.File, qtdMaxElem int) {
	// nao passar de qtdMaxElem para os dois vetores de input e para o vetor de output
	// sempre que a ouput encher, escreve no arquivo
	// ordenar
	// escrever
}

func Merge_Files(readData func(file *os.File, num int) []util.T, sortAlg string, size int, memMax int) {
	file, err := os.Open("integerscpp.bin")
	if err != nil {
		log.Fatal("Erro na leitura do arquivo binario", err)
		defer file.Close()
	}

	stat, _ := file.Stat()
	stat.Size()

	dataNumber := int(math.Floor(math.Pow(2, 30) / float64(size)))
	// fileLimit := size*dataNumber

	readData(file, dataNumber)

	// switch sortAlg {
	// 	case "Merge":
	// 		go Mergesort()
	// 	case "Quick":
	// 		go Quicksort()
	// }

	fmt.Println(readData(file, 2))
	fmt.Println(readData(file, 4))
}
