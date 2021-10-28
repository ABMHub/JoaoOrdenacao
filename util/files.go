package util

import (
	//"bytes"
	"encoding/binary"
	//"file"
	"fmt"
	"io"
	"math/rand"

	//"unsafe"

	//"log"
	"os"
)

type ReadData func(file *os.File, num int64) []T
type Compare func(T, T) bool
type WriteData func(file *os.File, array []T)
type Fragment_files func(file_name string, elem_size int, max_size int64) ([]*os.File, []int)

func FragmentBin(file_name string, elem_size int, max_size int64) ([]*os.File, []int) {
	//Abre o arquivo para descobrir o tamanho
	file, _ := os.Open(file_name)
	//Normaliza max_size
	max_size -= max_size % int64(elem_size)

	//Descobre o tamanho do arquivo
	stat, _ := file.Stat()
	size := stat.Size()

	file.Close()

	//Define a quantidade de ponteiros para arquivo de acordo com o tamanho do elemento
	fds_qtd := size / max_size
	qtd_elem := max_size / int64(elem_size)

	//Define um ponteiro de fds
	fds := make([]*os.File, fds_qtd)
	//Define o tamanho de conteudo dos file descriptors
	size_fd := make([]int, fds_qtd)

	for i := 0; i < int(fds_qtd); i++ {
		// obtem os fds
		fds[i], _ = os.Open(file_name)
		fds[i].Seek(max_size*int64(i), 0)
		
		// obtem os tamanhos
		size_fd[i] = int(qtd_elem)
	}

	return fds, size_fd
}

//Recebe o arquivo a ser lido e o tamanho em bytes do elemento que deve ser lido
func ReadBytes(file *os.File, qtdBytes int) ([]byte, error) {
	bytes := make([]byte, qtdBytes) //Vetor com os bytes do elemento

	_, err := file.Read(bytes)       //Le do arquivo
	if err != nil && err != io.EOF { //Se encontrar um erro, printa uma mensagem de erro
		fmt.Println("ReadBytes Linha 16, Erro ao ler o arquivo binário")
		//log.Fatal(err)
	}

	//Retorna os bytes do elemento e o err de leitura
	return bytes, err
}

type Pair3 struct {
	Fst, Snd int32
}

type lel struct {
	ff int
}

// func WriteIntegers(file *os.File, arr []T) {
// 	//buf := new(bytes.Buffer)

// 	for i := 0; i < len(arr); i++ {
// 		err := binary.Write(file, binary.LittleEndian, arr[i].(uint32))
// 		if err != nil {
// 			fmt.Println("binary.Write failed:", err)
// 		}
// 	}

// }

func WriteIntegers(file *os.File, arr []T) {
	t2 := make([]byte, len(arr)*4)

	//t := *(*[] Pair3)(unsafe.Pointer(&arr))
	j := 0
	for i := 0; i < len(t2); i += 4 {
		binary.LittleEndian.PutUint32(t2[i:i+4], arr[j].(uint32))
		j += 1
	}

	//for i := 0; i < len(arr); i++ {
	err := binary.Write(file, binary.LittleEndian, t2)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	//}
}

//Recebe o arquivo que sera lido e a quantidade de elementos a serem lidos
func ReadIntegers(file *os.File, num int64) []T {
	var arr []T
	var i int64
	for i = 0; i < num; i++ {
		//Pega o valor em bytes de um inteiro no arquivo e o erro que ocorreu
		bytes, err := ReadBytes(file, 4)

		//Se o erro indicar o EOF, dah um break no loop, dessa forma quando chegar no fim do arquivo
		//o tamanho do vetor arr vai ser menor que num
		if err == io.EOF {
			break
		}

		//Converte os bytes obtidos da ReadBytes no tipo T generico
		arr = append(arr, T(binary.LittleEndian.Uint32(bytes)))
	}

	//Retorna os vetor com os elementos da leitura
	return arr
}

func GenerateFiles(n int64) {
	ptr, err := os.Create("IntegersGo.bin")
	if err != nil {
		fmt.Println("erro")
	}

	var i int64
	for i = 0; i < n; i++ {
		err := binary.Write(ptr, binary.LittleEndian, uint32(rand.Int()))
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
}
