package sort

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"sortAlgorithms/util"
	"sync"

	"golang.org/x/sync/semaphore"
)

var queueLock = &sync.Mutex{}
var sem *(semaphore.Weighted)

func Merge_arrays(cmp func(util.T, util.T) bool, file1, file2 *os.File, qtdMaxElem int) {
	// nao passar de qtdMaxElem para os dois vetores de input e para o vetor de output
	// sempre que a ouput encher, escreve no arquivo
	// ordenar
	// escrever
}

func Read_And_Sort(page, elem_size int, fileLimit int64, file_name string, readData func(file *os.File, num int64) []util.T) {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal("Erro na leitura do arquivo binario", err)
		defer file.Close()
	}
	file.Seek(int64(page)*fileLimit, 0)
	readData(file, fileLimit/(int64(elem_size)))
	sem.Release(1)
}

func Merge_Files(readData func(file *os.File, num int64) []util.T, sortAlg string, size int, memMax int) {
	file, err := os.Open("integerscpp.bin") // abre arquivo
	if err != nil {                         // se der erro cancela tudo
		log.Fatal("Erro na leitura do arquivo binario", err) //
		defer file.Close()                                   //
	}

	stat, _ := file.Stat()
	stat.Size() // tamanho do arquivo

	dataNumber := int64(math.Floor(math.Pow(2, 30) / float64(size))) // qtd de file descriptors
	fileLimit := int64(size) * dataNumber                            // numero em bytes do offset do seek

	sem = semaphore.NewWeighted(8) // semaphoro com 8 permissoes
	ctx := context.Background()    // nao sei????????? (necessario pro sem.acquire)

	for i := 0; i < dataNumber; i++ {
		sem.Acquire(ctx, 1) // pega uma permissao do sem
		Read_And_Sort(i, fileLimit, "integerscpp.bin")
	}

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
