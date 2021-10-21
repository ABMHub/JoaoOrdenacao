package main

import (
	// "fmt"
	"log"
	"os"

	"fmt"
	// "time"

	"sortAlgorithms/sort"
	"sortAlgorithms/util"
	// "github.com/cheggaaa/pb/v3"
)

func main() {
	// contador := 100000
	// barra := pb.StartNew(contador)

	// for i := 0; i < contador; i++ {
	// 	barra.Increment()
	// 	time.Sleep(time.Millisecond)
	// }
	size := 4
	util.GenerateFiles(1000000)
	sort.Merge_Files(util.ReadIntegers, "merge-sort", size, util.CompareInt)
	// util.SetThreadLimit(1)

	file1, err1 := os.Open("IntegersGo.bin") // abre arquivo
	if err1 != nil {                         // se der erro cancela tudo
		log.Fatal("Erro na leitura do arquivo binario com os inteiros a serem ordenados", err1) //
		defer file1.Close()                                                                     //
	}

	fmt.Println(util.ReadIntegers(file1, 10))
	fmt.Println("oi")
}
