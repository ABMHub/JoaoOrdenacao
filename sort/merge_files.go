package sort

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
	"sortAlgorithms/util"
	"sync"
)

var queueLock = &sync.Mutex{}

func Merge_arrays(readData func(file *os.File, num int) []util.T, cmp func(util.T, util.T) bool, file1, file2 *os.File, qtdMaxElem int) {

	//Cria o arquivo com o output
	fileO, err := os.Create("output.bin")
	if err != nil {
		log.Fatal(err)
	}

	defer fileO.Close()

	var idx int = 0                     //Indice do vetor de output
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
				err = binary.Write(fileO, binary.LittleEndian, outArr)
				if err != nil {
					fmt.Println("Nao foi possivel escrever no arquivo output", err)
				}
				outArr = nil //Zera o vetor
				idx = 0      //O vetor output tem 0 elementos
			}
		}
	}

	if len(outArr) != 0 { //Se o vetor de output não for vazio, escreve no arquivo de output
		//Escreve os dados no arquivo output
		err = binary.Write(fileO, binary.LittleEndian, outArr)
		if err != nil {
			fmt.Println("Nao foi possivel escrever no arquivo output", err)
		}
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
		err = binary.Write(fileO, binary.LittleEndian, inArr1)
		if err != nil {
			fmt.Println("Nao foi possivel escrever no arquivo output", err)
		}
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
		err = binary.Write(fileO, binary.LittleEndian, inArr2)
		if err != nil {
			fmt.Println("Nao foi possivel escrever no arquivo output", err)
		}
		inArr2 = nil //Zera o vetor do arquivo 2
	}
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
