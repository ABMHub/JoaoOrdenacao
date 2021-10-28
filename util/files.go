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

	"github.com/cheggaaa/pb"
)

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
	// fmt.Println("adasdasds")
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

func GenerateFiles(n int64, random_seed int64) {
	progress_bar := pb.New64(n)
	progress_bar.Prefix("Generate Files")
	progress_bar.ShowSpeed = false
	progress_bar.ShowElapsedTime = true
	progress_bar.Start()

	ptr, err := os.Create("IntegersGo.bin")

	if err != nil {
		fmt.Println("erro")
	}

	m := n / 10000

	rand.Seed(random_seed)

	var i int64
	var j int64
	a := make([]T, m)
	for i < n {
		for j = 0; j < m && i < n; j++ {
			i++
			a[j] = uint32(rand.Int())
			progress_bar.Increment()
		}
		WriteIntegers(ptr, a[0:j])
	}
	progress_bar.Finish()
}
