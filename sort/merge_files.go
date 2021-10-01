package sort

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"sortAlgorithms/util"
	"strconv"

	//"sync"

	"golang.org/x/sync/semaphore"
)

//var queueLock = &sync.Mutex{}
var sem *(semaphore.Weighted)

func Merge_arrays(readData func(file *os.File, num int64) []util.T, cmp func(util.T, util.T) bool, file1, file2 *os.File, qtdMaxElem int64) {

	//Cria o arquivo com o output
	fileO, err := os.Create("output.bin")
	if err != nil {
		log.Fatal(err)
	}

	defer fileO.Close()

	var idx int64 = 0                   //Indice do vetor de output
	var flag1, flag2 bool = true, true  //Flags para indicar se ainda ha elementos nos arquivos
	var outArr, inArr1, inArr2 []util.T //Vetores que conterao os elementos lidos dos arquivos

	for flag1 && flag2 { //Enquanto houver elementos no arquivo 1 e arquivo 2
		if len(inArr1) == 0 { //Se o vetor do arquivo 1 for vazio, pega os elementos do arquivo
			inArr1 = readData(file1, qtdMaxElem/4) //Pega qtdMaxElem/4 elementos do arquivo

			//Se ainda assim o vetor do arquivo 1 for vazio, sai do loop (acabou os elementos)
			if len(inArr1) == 0 {
				flag1 = false //Indica que nao tem mais elementos para ler do arquivo 1
				break
			}
		}
		if len(inArr2) == 0 { //Se o vetor do arquivo 2 for vazio, pega os elementos do arquivo
			inArr2 = readData(file2, qtdMaxElem/4) //Pega qtdMaxElem/4 elementos do arquivo

			//Se ainda assim o vetor do arquivo 2 for vazio, sai do loop (acabou os elementos)
			if len(inArr2) == 0 {
				flag2 = false //Indica que nao tem mais elementos para ler do arquivo 2
				break
			}
		}

		for len(inArr1) != 0 && len(inArr2) != 0 { //Enquanto houver elementos nos vetores
			if cmp(inArr1[0], inArr2[0]) { //Se retornar true, entao inArr1[0] < inArr2[0]
				outArr = append(outArr, inArr1[0]) //Adiciona o elemento inArr1[0] no vetor de output
				inArr1 = inArr1[1:]                //Remove o primeiro elemento do vetor inArr1
			} else { //Se retornar true, entao inArr1[0] >= inArr2[0]
				outArr = append(outArr, inArr2[0]) //Adiciona o elemento inArr2[0] no vetor de output
				inArr2 = inArr2[1:]                //Remove o primeiro elemento do vetor inArr2
			}

			idx++ //Aumentou em 1 a quantidade de elementos no vetor de output

			if idx == qtdMaxElem/2 { //Se o vetor de output estiver cheio
				//Escreve os dados no arquivo output
				util.WriteIntegers(fileO, outArr)
				outArr = nil //Zera o vetor
				idx = 0      //O vetor output tem 0 elementos
			}
		}
	}

	if len(outArr) != 0 { //Se o vetor de output não for vazio, escreve no arquivo de output
		//Escreve os dados no arquivo output
		util.WriteIntegers(fileO, outArr)
	}

	for flag1 { //Se ainda houver elementos no arquivo 1
		if len(inArr1) == 0 { //Se o vetor do arquivo 1 for vazio, pega os elementos do arquivo
			inArr1 = readData(file1, qtdMaxElem/4) //Pega qtdMaxElem/4 elementos do arquivo

			//Se ainda assim o vetor do arquivo 1 for vazio, sai do loop
			if len(inArr1) == 0 {
				break
			}
		}

		//Escreve os dados no arquivo output
		util.WriteIntegers(fileO, inArr1)
		inArr1 = nil //Zera o vetor do arquivo 1
	}

	for flag2 { //Se ainda houver elementos no arquivo 2
		if len(inArr2) == 0 { //Se ainda houver elementos no arquivo 1
			inArr2 = readData(file2, qtdMaxElem/4) //Pega qtdMaxElem/4 elementos do arquivo

			//Se ainda assim o vetor do arquivo 2 for vazio, sai do loop
			if len(inArr2) == 0 {
				break
			}
		}

		//Escreve os dados no arquivo output
		util.WriteIntegers(fileO, inArr2)
		inArr2 = nil //Zera o vetor do arquivo 2
	}
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
	// err = os.Mkdir("temp", 0755)
	// if err != nil {
	//  	fmt.Println("A pasta ja existia", err)
	// }

	os.Mkdir("temp", 0755)
	//Cria um arquivo temporario
	fout, err := os.Create("temp" + string(os.PathSeparator) + "out" + strconv.Itoa(page) + ".bin")
	if err != nil {
		fmt.Println("Nao foi possivel criar o arquivo temporario"+strconv.Itoa(page), err)
		//return
	}

	util.WriteIntegers(fout, arr)
	// codigo experimental
	// var buf bytes.Buffer
	// enc := gob.NewEncoder(&buf)
	// enc.Encode(arr)

	if err != nil {
		fmt.Println("Nao foi possivel escrever no arquivo temporario"+strconv.Itoa(page), err)
		//fout.Close()
		//return
	}

	//fim do codigo experimental

	//Escreve os dados no arquivo temporario
	// err = binary.Write(fout, binary.LittleEndian, arr)
	// if err != nil {
	// 	fmt.Println("Nao foi possivel escrever no arquivo temporario" + strconv.Itoa(page), err)
	// 	//fout.Close()
	// 	//return
	// }

	//Fecha o arquivo e libera uma posicao no semaforo
	fout.Close()
	sem.Release(1)
}

func Merge_Files(readData func(file *os.File, num int64) []util.T, sortAlg string, size int, memMax int, cmp func(util.T, util.T) bool) {
	file, err := os.Open("integerscpp2.bin") // abre arquivo
	if err != nil {                          // se der erro cancela tudo
		log.Fatal("Erro na leitura do arquivo binario com os inteiros a serem ordenados", err) //
		defer file.Close()                                                                     //
	}

	stat, _ := file.Stat()
	//stat.Size() // tamanho do arquivo
	unidade := 3
	//dataNumber := int(math.Floor(math.Pow(2, float64(unidade)) / float64(size))) * 10 // qtd de file descriptors
	fds_qtd := int(math.Floor(float64(stat.Size())/math.Pow(10, float64(unidade)))) / size
	file_limit := stat.Size() / int64(fds_qtd)
	//fileLimit := int64(size * dataNumber)                          // numero em bytes do offset do seek
	fmt.Println(file_limit)
	sem = semaphore.NewWeighted(8) // semaphoro com 8 permissoes
	ctx := context.Background()    // nao sei????????? (necessario pro sem.acquire)

	for i := 0; i < fds_qtd; i++ {
		sem.Acquire(ctx, 1) // pega uma permissao do sem
		Read_And_Sort(i, size, file_limit, "integerscpp2.bin", sortAlg, readData, cmp)
	}

}
