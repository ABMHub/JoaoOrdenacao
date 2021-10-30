package main

import (
	"fmt"
	"log"
	"os"

	"sortAlgorithms/sort"
	"sortAlgorithms/util"
	// "time"
	// "github.com/cheggaaa/pb"
)

func main() {
	size := 4

	sort.Merge_Files("IntegersGo.bin", "merge-sort", size, 1, util.ReadIntegers, util.CompareInt, util.FragmentBin, util.WriteIntegers)
	util.SetThreadLimit(1)
	//util.GenerateFiles(250000000)

	file1, err1 := os.Open("Sorted.bin") // abre arquivo
	if err1 != nil {                     // se der erro cancela tudo
		log.Fatal("Erro na leitura do arquivo binario com os inteiros a serem ordenados", err1) //
		defer file1.Close()                                                                     //
	}

	fmt.Println(util.ReadIntegers(file1, 10))
	fmt.Println("oi")

	// a := pb.New(10)
	// b := pb.New(10)
	// c := pb.NewPool(a, b)

	// screen.Clear()
	// c.Start()

	// for i := 0; i < 3; i++ {
	// 	a.Increment()
	// 	time.Sleep(1000000000)
	// 	b.Increment()
	// 	time.Sleep(1000000000)
	// }

	// d := pb.New(6)
	// c.Add(d)

	// for i := 0; i < 3; i++ {
	// 	a.Increment()
	// 	time.Sleep(1000000000)
	// 	b.Increment()
	// 	time.Sleep(1000000000)
	// 	d.Increment()
	// 	time.Sleep(1000000000)
	// }

	// c.Stop()
	// screen.Clear()
	// c = pb.NewPool(d, a)
	// c.Start()

	// for i := 0; i < 4; i++ {
	// 	a.Increment()
	// 	time.Sleep(1000000000)
	// 	d.Increment()
	// 	time.Sleep(1000000000)
	// }

	// a.Finish()
	// b.Finish()

	// pbar := pb.New(6)
	// ppool := util.NewPBar(pbar)
	// for i := 0; i < 3; i++ {
	// 	pbar.Increment()
	// 	time.Sleep(1000000000)
	// }
	// pbar2 := pb.New(6)
	// ppool.Add(pbar2)
	// for i := 0; i < 3; i++ {
	// 	pbar.Increment()
	// 	pbar2.Increment()
	// 	time.Sleep(1000000000)
	// }
	// pbar.Finish()
	// ppool.UpdateFinished()
	// for i := 0; i < 3; i++ {
	// 	pbar2.Increment()
	// 	time.Sleep(1000000000)
	// }
	// pbar2.Finish()
	// ppool.End()
}
