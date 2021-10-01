package util

import (
	"encoding/binary"
	"fmt"
	"io"
	"bytes"
	"log"
	"os"
)

var network bytes.Buffer

//Recebe o arquivo a ser lido e o tamanho em bytes do elemento que deve ser lido
func ReadBytes(file *os.File, qtdBytes int) ([]byte, error) {
	bytes := make([]byte, qtdBytes) //Vetor com os bytes do elemento

	_, err := file.Read(bytes) //Le do arquivo
	if err != nil && err != io.EOF {            //Se encontrar um erro, printa uma mensagem de erro
		fmt.Println("ReadBytes Linha 16, Erro ao ler o arquivo binário")
		log.Fatal(err)
	}

	//Retorna os bytes do elemento e o err de leitura
	return bytes, err
}

func WriteIntegers(file *os.File, arr []T) {
	//buf := new(bytes.Buffer)

	for i := 0; i < len(arr); i++ {
		err := binary.Write(file, binary.LittleEndian, arr[i].(uint32))
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
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
