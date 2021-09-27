package sort

import (
	"encoding/binary"
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"sortAlgorithms/util"
	"strconv"
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

func Read_And_Sort(page, elem_size int, fileLimit int64, file_name, sortAlg string, readData func(file *os.File, num int64) []util.T, cmp func(util.T, util.T) bool) {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal("Erro na leitura do arquivo binario que sera ordenado", err)
		defer file.Close()
	}

	file.Seek(int64(page)*fileLimit, 0)                 //posiciona o ponteiro onde o arquivo deve ser lido
	arr := readData(file, fileLimit/(int64(elem_size))) //le os dados a partir da posicao definida

	//ordena os dados lidos
	switch sortAlg {
	case "quick-sort":
		Quicksort_F(arr, 0, len(arr)-1, cmp)
	case "merge-sort":
		Mergesort_F(arr, 0, len(arr)-1, cmp)
	default:
		Mergesort_F(arr, 0, len(arr)-1, cmp)
	}

	//Cria um diretorio onde serao salvos os arquivos temporarios caso ele ainda nao exista
	//err = os.Mkdir("temp", 0755)
	// if err != nil {
	// 	fmt.Println("A pasta ja existia", err)
	// }
	os.Mkdir("temp", 0755)
	
	//Cria um arquivo temporario
	fout, err := os.Create("temp" + string(os.PathSeparator) + "out" + strconv.Itoa(page) + ".bin")
	if err != nil {
		fmt.Println("Nao foi possivel criar o arquivo temporario" + strconv.Itoa(page), err)
		//return
	}

	//Escreve os dados no arquivo temporario
	err = binary.Write(fout, binary.LittleEndian, arr)
	if err != nil {
		fmt.Println("Nao foi possivel escrever no arquivo temporario" + strconv.Itoa(page), err)
		//fout.Close()
		//return
	}

	//Fecha o arquivo e libera uma posicao no semaforo
	fout.Close()
	sem.Release(1)
}

func Merge_Files(readData func(file *os.File, num int64) []util.T, sortAlg string, size int, memMax int, cmp func(util.T, util.T) bool) {
	// file, err := os.Open("integerscpp.bin") // abre arquivo
	// if err != nil {                         // se der erro cancela tudo
	// 	log.Fatal("Erro na leitura do arquivo binario com os inteiros a serem ordenados", err) //
	// 	defer file.Close()                                   //
	// }

	//stat, _ := file.Stat()
	//stat.Size() // tamanho do arquivo
	unidade := 10.0
	dataNumber := int(math.Floor(math.Pow(2, unidade) / float64(size))) * 10 // qtd de file descriptors
	fileLimit := int64(size * dataNumber)                          // numero em bytes do offset do seek

	sem = semaphore.NewWeighted(8) // semaphoro com 8 permissoes
	ctx := context.Background()    // nao sei????????? (necessario pro sem.acquire)

	for i := 0; i < dataNumber; i++ {
		sem.Acquire(ctx, 1) // pega uma permissao do sem
		Read_And_Sort(i, size, fileLimit, "integerscpp.bin", sortAlg, readData, cmp)
	}

}
