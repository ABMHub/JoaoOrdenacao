package main

import (
	"time"

	"github.com/cheggaaa/pb"
)

func main() {
	// size := 4
	// // util.GenerateFiles(250000000)
	// sort.Merge_Files(util.ReadIntegers, "merge-sort", size, util.CompareInt)
	// // util.SetThreadLimit(1)

	// file1, err1 := os.Open("IntegersGo.bin") // abre arquivo
	// if err1 != nil {                         // se der erro cancela tudo
	// 	log.Fatal("Erro na leitura do arquivo binario com os inteiros a serem ordenados", err1) //
	// 	defer file1.Close()                                                                     //
	// }

	// fmt.Println(util.ReadIntegers(file1, 10))
	// fmt.Println("oi")

	// todo pool global
	// todo sempre que matar uma pb, mata a pool -> clear -> revive a pool

	a := pb.New(10)
	b := pb.New(10)
	c := pb.NewPool(a, b)

	c.Start()

	for i := 0; i < 3; i++ {
		a.Increment()
		time.Sleep(1000000000)
		b.Increment()
		time.Sleep(1000000000)
	}

	d := pb.New(6)
	c.Add(d)

	for i := 0; i < 3; i++ {
		a.Increment()
		time.Sleep(1000000000)
		b.Increment()
		time.Sleep(1000000000)
		d.Increment()
		time.Sleep(1000000000)
	}

	c.Stop()
	c = pb.NewPool(a, d)
	c.Start()

	for i := 0; i < 4; i++ {
		a.Increment()
		time.Sleep(1000000000)
		d.Increment()
		time.Sleep(1000000000)
	}

	a.Finish()
	b.Finish()
}
