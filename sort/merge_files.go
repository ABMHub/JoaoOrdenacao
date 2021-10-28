package sort

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"sortAlgorithms/util"
	"strconv"
	"sync"

	//"sync"

	"golang.org/x/sync/semaphore"
)

var wg sync.WaitGroup

var cond_files = &sync.Cond{} //variavel condicao para os arquivos

var queueLock = &sync.Mutex{}

const sem_permissions_RAS int = 12 //numero de permissoes que o semaforo Read_And_Sort
var sem_RAS *(semaphore.Weighted)  //controla as threads Read_And_Sort

//var sem_files *(semaphore.Weighted)		//semaforo que controla os arquivos que ja podem ser mesclados

var count_files int //quantidade de arquivos temporarios
//Fila com os arquivos prontos
var files_queue util.List

func Merge_arrays(readData util.ReadData, cmp util.Compare,
	writeData util.WriteData, file1_n, file2_n string, qtdMaxElem int64, output_name string) {
	file1, err1 := os.Open("temp" + string(os.PathSeparator) + file1_n + ".bin") // abre arquivo
	if err1 != nil {
		log.Fatal(err1)
	}

	file2, err2 := os.Open("temp" + string(os.PathSeparator) + file2_n + ".bin") // abre arquivo
	if err2 != nil {
		log.Fatal(err2)
	}

	//Cria o arquivo com o output
	folder := "temp" + string(os.PathSeparator)
	//path := output_name + ".bin"
	fileO, err := os.Create(folder + output_name + ".bin")
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
				writeData(fileO, outArr)
				outArr = nil //Zera o vetor
				idx = 0      //O vetor output tem 0 elementos
			}
		}
	}

	if len(outArr) != 0 { //Se o vetor de output nÃ£o for vazio, escreve no arquivo de output
		//Escreve os dados no arquivo output
		writeData(fileO, outArr)
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
		writeData(fileO, inArr1)
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
		writeData(fileO, inArr2)
		inArr2 = nil //Zera o vetor do arquivo 2
	}

	//queueLock.Lock()
	cond_files.L.Lock()
	files_queue.Push_back(output_name)
	//sem_files.Release(1)			//Significa que um arquivo foi adicionado na fila
	count_files += 1 //Incrementa o numero de arquivos prontos
	if count_files >= 2 {
		cond_files.Signal()
	}
	cond_files.L.Unlock()
	//queueLock.Unlock()

	//Fecha os arquivos
	file1.Close()
	file2.Close()

	//Deleta os arquivos que foram mesclados
	fmt.Println("f1: ", file1_n, "f2:", file2_n)
	os.Remove("temp" + string(os.PathSeparator) + file1_n + ".bin") // deleta arquivo
	os.Remove("temp" + string(os.PathSeparator) + file2_n + ".bin") // deleta arquivo

	sem_RAS.Release(1)
	wg.Done() //Sinaliza que a thread acabou
}

/*
	Recebe como parametro um arquivo e um indice (page) a partir de qual parte desse arquivo
	deve ordenar
*/
func Read_And_Sort(file_name, sort_alg string, fds *os.File, readData util.ReadData, cmp util.Compare) {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal("Erro na leitura do arquivo binario que sera ordenado", err)
		defer file.Close()
	}

	//ordena os dados lidos
	switch sort_alg {
	case "quick-sort":
		Quicksort_F(arr, 0, len(arr)-1, cmp)
	case "merge-sort":
		Mergesort_P(arr, cmp)
	default:
		Mergesort_P(arr, cmp)
	}

	//Cria uma pasta temporaria
	os.Mkdir("temp", 0755)

	//define o nome do arquivo e sua path
	folder := "temp" + string(os.PathSeparator)
	path := "out" + strconv.Itoa(page)

	//Cria um arquivo temporario
	fout, err := os.Create(folder + path + ".bin")
	if err != nil {
		fmt.Println("Nao foi possivel criar o arquivo temporario"+strconv.Itoa(page), err)
	}

	util.WriteIntegers(fout, arr)

	//Coloca a path do arquivo ordenado, lock porque eh regiao critica
	//queueLock.Lock()
	cond_files.L.Lock()
	files_queue.Push_back(path)
	//sem_files.Release(1)			//Significa que um arquivo foi adicionado na fila
	count_files += 1 //Incrementa o numero de arquivos prontos
	if count_files >= 2 {
		cond_files.Signal()
	}
	cond_files.L.Unlock()
	//queueLock.Unlock()

	if err != nil {
		fmt.Println("Nao foi possivel escrever no arquivo temporario"+strconv.Itoa(page), err)
	}

	fout.Close()
	sem_RAS.Release(1)
}

// max_size eh em MB
func Merge_Files(file_name, sortAlg string, max_size int, readData util.ReadData, cmp util.Compare, fragment util.Fragment_files) {
	// unidade de medida para max_size 
	const size_unity = 1000000

	// define a quantidade maxima de memoria
	max_size = max_size * size_unity
	
	// abre arquivo
	file, err := os.Open(file_name) 
	if err != nil {               // se der erro cancela tudo
		log.Fatal("Erro na leitura do arquivo binario com os inteiros a serem ordenados", err) //
		defer file.Close()                                                                     //
	}

	//Inicializa a variavel condicao
	cond_files = sync.NewCond(queueLock)

	//Inicializa a fila que vai conter os arquivos ja ordenados
	files_queue = util.NewList()
	count_files = 0

	//Contexto da thread principal
	ctx := context.Background()

	//semaforo que controla as threads da read and sort
	sem_RAS = semaphore.NewWeighted(int64(sem_permissions_RAS))

	// Obtem uma lista com os file descriptors necessario
	fds := fragment(file, max_size)	
	fds_qtd := len(fds)

	//fragmenta e ordena os arquivos
	var i int
	for i = 0; i < fds_qtd; i++ {
		sem_RAS.Acquire(ctx, 1) // pega uma permissao do sem
		go Read_And_Sort(file_name, sortAlg, fds[i], readData, cmp)
	}

	var output_name string
	//controla a mesclagem de arquivos
	for count := 0; count < fds_qtd-1; count += 1 {

		sem_RAS.Acquire(ctx, 1) //Garante que o numero de threads esteja dentro do permitido

		//Chama o procedimento que mescla os arquivos

		//So acontece quando pelo menos dois arquivos estiverem prontos
		cond_files.L.Lock()

		for count_files < 2 {
			cond_files.Wait()
		}
		file1_name := (files_queue.Pop_front()).(string)
		file2_name := (files_queue.Pop_front()).(string)
		count_files -= 2

		cond_files.L.Unlock()

		output_name = "out" + strconv.Itoa(i)
		i++
		wg.Add(1)
		go Merge_arrays(util.ReadIntegers, util.CompareInt, util.WriteIntegers, file1_name, file2_name, int64(n_max_elements), output_name)
	}

	wg.Wait() //Espera todo mundo terminar

	//Renomeia o arquivo, move ele pra raiz e deleta a temp
	os.Rename("temp"+string(os.PathSeparator)+output_name+".bin", "."+string(os.PathSeparator)+"Sorted"+".bin")
	os.Remove("temp")
}
