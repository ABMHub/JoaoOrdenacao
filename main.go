package main

import (
	// "fmt"
	// "log"
	 "os"

	"fmt"

	"sortAlgorithms/sort"
	"sortAlgorithms/util"
	"encoding/binary"
	"encoding/json"
	"time"
)

func main() {


	file,_ := os.Open("ints.txt")
	util.SetThreadLimit(3)
	temp := make([]byte, 4)
	arr := make([]util.T,1e6)
	for i := 0; i < 1e6; i++{
		file.Read(temp)
		arr[i] = binary.LittleEndian.Uint32(temp)
	}
	start := time.Now()
	fmt.Println(len(arr)-1)
	sort.Quicksort_P(arr,0,len(arr)-1,func(a, b util.T)bool{
		return a.(uint32) > b.(uint32)
	})
	fmt.Println(time.Since(start))
	fmt.Println(util.IsSorted(arr,func(a, b util.T)bool{
		return a.(uint32) > b.(uint32)
	}))


	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"},
		{"Name": "Bitchass", "Order": "bac"}]`)
	fmt.Println(len(jsonBlob))
	p := make(map[string]util.T)
	err := json.Unmarshal(jsonBlob, &p)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", p)
	//util.GenerateFiles(250000000)

	// file1, err1 := os.Open("IntegersGo.bin") // abre arquivo
	// if err1 != nil {                         // se der erro cancela tudo
	// 	log.Fatal("Erro na leitura do arquivo binario com os inteiros a serem ordenados", err1) //
	// 	defer file1.Close()                                                                     //
	// }

	//fmt.Println(util.ReadIntegers(file1, 10))
	fmt.Println("oi")
}
