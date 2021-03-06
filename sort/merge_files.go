package sort

import (
	"context"
	"fmt"
	"log"
	"os"
	"sortAlgorithms/util"
	"strconv"
	"sync"
	"time"

	"github.com/cheggaaa/pb"
	"golang.org/x/sync/semaphore"
)

var pPool *util.PBar
var wg sync.WaitGroup
var general_pbar *pb.ProgressBar

var cond_files = &sync.Cond{} //variavel condicao para os arquivos

var queueLock = &sync.Mutex{}
var poolLock = &sync.Mutex{}

var sem_RAS *(semaphore.Weighted) //controla as threads Read_And_Sort

//var sem_files *(semaphore.Weighted)		//semaforo que controla os arquivos que ja podem ser mesclados

var count_files int //quantidade de arquivos temporarios
//Fila com os arquivos prontos
var files_queue util.List

func merge_arrays(file1_n, file2_n, output_name string, elem_size int, max_size int64, readData util.ReadData,
	writeData util.WriteData, cmp util.Compare) {

	file1, err1 := os.Open("temp" + string(os.PathSeparator) + file1_n + ".bin") // abre arquivo
	if err1 != nil {
		log.Fatal(err1)
	}

	file2, err2 := os.Open("temp" + string(os.PathSeparator) + file2_n + ".bin") // abre arquivo
	if err2 != nil {
		log.Fatal(err2)
	}

	stat1, _ := file1.Stat()
	stat2, _ := file2.Stat()

	merge_progress_bar := pb.New64((stat1.Size() + stat2.Size()) / int64(elem_size))
	merge_progress_bar.Prefix(stat1.Name() + " + " + stat2.Name())
	merge_progress_bar.ShowSpeed = false
	// merge_progress_bar.ShowElapsedTime = true
	poolLock.Lock()
	pPool.Add(merge_progress_bar)
	poolLock.Unlock()

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
			inArr1 = readData(file1, max_size/4) //Pega max_size/4 elementos do arquivo

			//Se ainda assim o vetor do arquivo 1 for vazio, sai do loop (acabou os elementos)
			if len(inArr1) == 0 {
				flag1 = false //Indica que nao tem mais elementos para ler do arquivo 1
				inArr1 = nil
				break
			}
		}
		if len(inArr2) == 0 { //Se o vetor do arquivo 2 for vazio, pega os elementos do arquivo
			inArr2 = readData(file2, max_size/4) //Pega max_size/4 elementos do arquivo

			//Se ainda assim o vetor do arquivo 2 for vazio, sai do loop (acabou os elementos)
			if len(inArr2) == 0 {
				flag2 = false //Indica que nao tem mais elementos para ler do arquivo 2
				inArr2 = nil
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
			merge_progress_bar.Increment()

			if idx == max_size/2 { //Se o vetor de output estiver cheio
				//Escreve os dados no arquivo output
				writeData(fileO, outArr)
				outArr = nil //Zera o vetor
				idx = 0      //O vetor output tem 0 elementos
			}
		}
	}

	if len(outArr) != 0 { //Se o vetor de output n????o for vazio, escreve no arquivo de output
		//Escreve os dados no arquivo output
		writeData(fileO, outArr)
	}

	for flag1 { //Se ainda houver elementos no arquivo 1
		if len(inArr1) == 0 { //Se o vetor do arquivo 1 for vazio, pega os elementos do arquivo
			inArr1 = readData(file1, max_size/4) //Pega max_size/4 elementos do arquivo

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
			inArr2 = readData(file2, max_size/4) //Pega max_size/4 elementos do arquivo

			//Se ainda assim o vetor do arquivo 2 for vazio, sai do loop
			if len(inArr2) == 0 {
				break
			}
		}

		//Escreve os dados no arquivo output
		writeData(fileO, inArr2)
		inArr2 = nil //Zera o vetor do arquivo 2
	}

	cond_files.L.Lock()							// Inicia uma regiao de exclusao mutua
	files_queue.Push_back(output_name)			// Adiciona o arquivo gerado pela mescla na fila
	count_files += 1 							// Incrementa o numero de arquivos prontos
	
	if count_files >= 2 {						// Se existirem ao menos dois arquivos prontos, acorda um processo 
		cond_files.Signal()						// que possa estar parado na merge_files esperando para que haja arquivos para 
	}											// mesclagem
	cond_files.L.Unlock()						// Libera a regiao de exclusao mutua

	//Fecha os arquivos
	file1.Close()
	file2.Close()

	//Deleta os arquivos que foram mesclados
	// fmt.Println("f1: ", file1_n, "f2:", file2_n)
	os.Remove("temp" + string(os.PathSeparator) + file1_n + ".bin") // deleta arquivo
	os.Remove("temp" + string(os.PathSeparator) + file2_n + ".bin") // deleta arquivo

	general_pbar.Increment()
	merge_progress_bar.Finish()
	poolLock.Lock()
	pPool.UpdateFinished()
	poolLock.Unlock()
	sem_RAS.Release(1)
	wg.Done() //Sinaliza que a thread acabou
}

/*
	Recebe como parametro um arquivo e um indice (page) a partir de qual parte desse arquivo
	deve ordenar
*/
func read_And_Sort(sort_alg string, page int, num_elem int64, fds util.T, readData util.ReadData, writeData util.WriteData, cmp util.Compare) {
	// le os dados
	arr := readData(fds, num_elem)

	//ordena os dados lidos
	switch sort_alg {
	case "quick-sort":
		Quicksort_P(arr, 0, len(arr)-1, cmp)
	case "merge-sort":
		Mergesort_P(arr, cmp)
	default:
		Mergesort_P(arr, cmp)
	}

	//Cria uma pasta temporaria se ela nao existir
	os.Mkdir("temp", 0755)

	//define o nome do arquivo e sua path
	folder := "temp" + string(os.PathSeparator)
	path := "out" + strconv.Itoa(page)

	//Cria um arquivo temporario
	fout, err := os.Create(folder + path + ".bin")
	if err != nil {
		fmt.Println("Nao foi possivel criar o arquivo temporario"+strconv.Itoa(page), err)
	}

	writeData(fout, arr)

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

	general_pbar.Increment()
	fout.Close()
	sem_RAS.Release(1)
	wg.Done()
}

// max_size eh em MB
func Merge_Files(file_name, sortAlg string, elemen_size, max_size, number_of_processors int, 
	readData util.ReadData, cmp util.Compare, fragment util.Fragment_files, writeData util.WriteData) {
	init_time := time.Now()						
	
	const size_unity = 1000000						// unidade de medida para max_size
	
	max_size = max_size * size_unity				// define a quantidade maxima de memoria
	
	cond_files = sync.NewCond(queueLock)			//Inicializa a variavel condicao
	
	files_queue = util.NewList()					//Inicializa a fila que vai conter os arquivos ja ordenados
	count_files = 0

	ctx := context.Background()						//Contexto da thread principal

	//semaforo que controla as threads da read and sort
	sem_RAS = semaphore.NewWeighted(int64(number_of_processors))

	// Obtem uma lista com os files descriptors necessarios
	fds, size_fd := fragment(file_name, number_of_processors, elemen_size, int64(max_size))
	fds_qtd := len(fds)

	// Controla a barra de progresso
	general_pbar = pb.New((fds_qtd * 2) - 1)
	general_pbar.Prefix("Total")
	general_pbar.ShowSpeed = false
	pPool = util.NewPBar(general_pbar)

	//fragmenta e ordena os arquivos
	var i int
	for i = 0; i < fds_qtd; i++ {
		wg.Add(1)
		sem_RAS.Acquire(ctx, 1) // pega uma permissao do sem
		go read_And_Sort(sortAlg, i, int64(size_fd[i]), fds[i], readData, writeData, cmp)
	}

	var output_name string
	//controla a mesclagem de arquivos

	for count := 0; count < fds_qtd-1; count += 1 {
		sem_RAS.Acquire(ctx, 1)								// Garante que o numero de threads esteja dentro do permitido
		
		cond_files.L.Lock()									// Da inicio a regiao de exclusao mutua

		for count_files < 2 {								// Caso o numero de arquivos prontos seja menor que 2, 
			cond_files.Wait()								// Aguarda na variavel condicao
		}		
		
		file1_name := (files_queue.Pop_front()).(string)	// Se chegou ate aqui, eh porque existem ao menos dois arquivos prontos
		file2_name := (files_queue.Pop_front()).(string)
		count_files -= 2									// Subtrai 2 da quantidade total de arquivos

		cond_files.L.Unlock()								// Da fim a regiao de exclusao mutua

		output_name = "out" + strconv.Itoa(i)
		i++
		wg.Add(1)
		go merge_arrays(file1_name, file2_name, output_name, elemen_size, int64(max_size), readData, writeData, cmp)
	}

	wg.Wait() //Espera todo mundo terminar

	if fds_qtd == 1 {
		output_name = "out" + strconv.Itoa(i-1)
	}

	//Renomeia o arquivo, move ele pra raiz e deleta a temp
	os.Rename("temp"+string(os.PathSeparator)+output_name+".bin", "."+string(os.PathSeparator)+"Sorted"+".bin")
	os.Remove("temp")

	pPool.End()
	fmt.Println(fds_qtd)
	since := time.Since(init_time)
	fmt.Printf("Tempo decorrido: %dm%ds\n", int(since.Minutes()), (int(since.Seconds()))%60)
}
