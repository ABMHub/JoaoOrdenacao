package test

import (
	// "fmt"
	"os"
	"sortAlgorithms/sort"
	"sortAlgorithms/util"
	"testing"
	// "time"
)

func TestCompare(t *testing.T) {
	x := []util.T{uint32(10), uint32(20), uint32(30), uint32(40), uint32(50)}
	v := []util.T{"lucas", "joao victor", "joao pedro", "alequi", "maria eduarda"}
	f := []util.T{float32(1.18), float32(2.72), float32(3.14), float32(4.57), float32(5.98)}

	if !util.CompareInt(x[0], x[1]) {
		t.Errorf("Erro no CompareInt, %d < %d eh falso", x[0], x[1])
	}

	if !util.CompareInt(x[2], x[3]) {
		t.Errorf("Erro no CompareInt, %d < %d eh falso", x[2], x[3])
	}

	if util.CompareInt(x[4], x[4]) {
		t.Errorf("Erro no CompareInt, %d < %d eh verdadeiro", x[4], x[4])
	}

	if !util.CompareString(v[3], v[1]) {
		t.Errorf("Erro no CompareString, %s < %s eh falso", v[3], v[1])
	}

	if !util.CompareString(v[2], v[1]) {
		t.Errorf("Erro no CompareString, %s < %s eh falso", v[2], v[1])
	}

	if util.CompareString(v[4], v[4]) {
		t.Errorf("Erro no CompareString, %s < %s eh verdadeiro", v[4], v[4])
	}

	if !util.CompareFloat(f[0], f[1]) {
		t.Errorf("Erro no CompareFloat, %f < %f eh falso", f[0], f[1])
	}

	if !util.CompareFloat(f[2], f[3]) {
		t.Errorf("Erro no CompareFloat, %f < %f eh falso", f[2], f[3])
	}

	if util.CompareFloat(f[4], f[4]) {
		t.Errorf("Erro no CompareFloat, %f < %f eh verdadeiro", f[4], f[4])
	}
}

func TestIsSorted(t *testing.T) {
	x := []util.T{uint32(10), uint32(20), uint32(30), uint32(40), uint32(50)}
	v := []util.T{"alequi", "joao pedro", "joao victor", "lucas", "maria eduarda"}
	f := []util.T{float32(1.18), float32(2.72), float32(3.14), float32(4.57), float32(5.98)}

	if !util.IsSorted(x, util.CompareInt) {
		t.Error("Erro, array eh indicado como nao ordenado")
	}

	if !util.IsSorted(v, util.CompareString) {
		t.Error("Erro, array [alequi, joao pedro, joao victor, lucas, maria eduarda] eh indicado como nao ordenado")
	}

	if !util.IsSorted(f, util.CompareFloat) {
		t.Error("Erro, array [1.18, 2.72, 3.14, 4.57, 5.98] eh indicado como nao ordenado")
	}

	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
		v[i], v[j] = v[j], v[i]
		f[i], f[j] = f[j], f[i]
	}

	if util.IsSorted(x, util.CompareInt) {
		t.Error("Erro, array [50, 40, 30, 20, 10] eh indicado como ordenado")
	}

	if util.IsSorted(v, util.CompareString) {
		t.Error("Erro, array [maria eduarda, lucas, joao victor, joao pedro, alequi] eh indicado como ordenado")
	}

	if util.IsSorted(f, util.CompareFloat) {
		t.Error("Erro, array [5.98, 4.57, 3.14, 2.72, 1.18] eh indicado como ordenado")
	}
}

func TestIsPerm(t *testing.T) {
	x1 := []util.T{uint32(10), uint32(20), uint32(30), uint32(40), uint32(50)}
	x2 := []util.T{uint32(20), uint32(40), uint32(50), uint32(30), uint32(10)}
	x3 := []util.T{uint32(60), uint32(40), uint32(10), uint32(30), uint32(10)}
	x4 := []util.T{uint32(10), uint32(20), uint32(30), uint32(40)}
	x5 := []util.T{uint32(10), uint32(20), uint32(30), uint32(40), uint32(50), uint32(60)}

	if !util.IsPerm(x1, x2) {
		t.Error("Erro, array [10, 20, 30, 40, 50] e [20, 40, 50, 30, 10] eh indicado como nao sendo permutacao")
	}

	if util.IsPerm(x1, x3) {
		t.Error("Erro, array [10, 20, 30, 40, 50] e [60, 40, 10, 30, 10] eh indicado como sendo permutacao")
	}

	if util.IsPerm(x1, x4) {
		t.Error("Erro, array [10, 20, 30, 40, 50] e [10, 20, 30, 40] eh indicado como sendo permutacao")
	}

	if util.IsPerm(x1, x5) {
		t.Error("Erro, array [10, 20, 30, 40, 50] e [10, 20, 30, 40, 50, 60] eh indicado como sendo permutacao")
	}

	v1 := []util.T{"alequi", "joao pedro", "joao victor", "lucas", "maria eduarda"}
	v2 := []util.T{"lucas", "joao victor", "joao pedro", "alequi", "maria eduarda"}
	v3 := []util.T{"pedro", "joao victor", "jose", "alequi", "maria eduarda"}
	v4 := []util.T{"alequi", "joao pedro", "joao victor", "lucas"}
	v5 := []util.T{"alequi", "joao pedro", "joao victor", "lucas", "maria eduarda", "pedro"}

	if !util.IsPerm(v1, v2) {
		t.Error("Erro, array [alequi, joao pedro, joao victor, lucas, maria eduarda] e [lucas, joao victor, joao pedro, alequi, maria eduarda] eh indicado como nao sendo permutacao")
	}

	if util.IsPerm(v1, v3) {
		t.Error("Erro, array [alequi, joao pedro, joao victor, lucas, maria eduarda] e [pedro, joao victor, jose, alequi, maria eduarda] eh indicado como sendo permutacao")
	}

	if util.IsPerm(v1, v4) {
		t.Error("Erro, array [alequi, joao pedro, joao victor, lucas, maria eduarda] e [alequi, joao pedro, joao victor, lucas] eh indicado como sendo permutacao")
	}

	if util.IsPerm(v1, v5) {
		t.Error("Erro, array [alequi, joao pedro, joao victor, lucas, maria eduarda] e [alequi, joao pedro, joao victor, lucas, maria eduarda, pedro] eh indicado como sendo permutacao")
	}

	f1 := []util.T{float32(1.18), float32(2.72), float32(3.14), float32(4.57), float32(5.98)}
	f2 := []util.T{float32(2.72), float32(4.57), float32(5.98), float32(3.14), float32(1.18)}
	f3 := []util.T{float32(3.69), float32(4.57), float32(7.89), float32(3.14), float32(1.18)}
	f4 := []util.T{float32(1.18), float32(2.72), float32(3.14), float32(4.57)}
	f5 := []util.T{float32(1.18), float32(2.72), float32(3.14), float32(4.57), float32(5.98), float32(7.89)}

	if !util.IsPerm(f1, f2) {
		t.Error("Erro, array [1.18, 2.72, 3.14, 4.57, 5.98] e [2.72, 4.57, 5.98, 3.14, 1.18] eh indicado como nao sendo permutacao")
	}

	if util.IsPerm(f1, f3) {
		t.Error("Erro, array [1.18, 2.72, 3.14, 4.57, 5.98] e [3.69, 4.57, 7.89, 3.14, 1.18] eh indicado como sendo permutacao")
	}

	if util.IsPerm(f1, f4) {
		t.Error("Erro, array [1.18, 2.72, 3.14, 4.57, 5.98] e [1.18, 2.72, 3.14, 4.57] eh indicado como sendo permutacao")
	}

	if util.IsPerm(f1, f5) {
		t.Error("Erro, array [1.18, 2.72, 3.14, 4.57, 5.98] e [1.18, 2.72, 3.14, 4.57, 5.98, 7.89] eh indicado como sendo permutacao")
	}
}

func TestBubbleSort(t *testing.T) {
	x := []util.T{uint32(20), uint32(40), uint32(50), uint32(30), uint32(10)}
	x1 := x
	v := []util.T{"lucas", "joao victor", "joao pedro", "alequi", "maria eduarda"}
	v1 := v
	f := []util.T{float32(2.72), float32(4.57), float32(5.98), float32(3.14), float32(1.18)}
	f1 := f

	sort.Bubblesort(x1, util.CompareInt)

	if !util.IsSorted(x1, util.CompareInt) || !util.IsPerm(x, x1) {
		t.Error("Bubblesort falhou em ordenar vetor de inteiros")
	}

	sort.Bubblesort(v1, util.CompareString)

	if !util.IsSorted(v1, util.CompareString) || !util.IsPerm(v, v1) {
		t.Error("Bubblesort falhou em ordenar vetor de strings")
	}

	sort.Bubblesort(f1, util.CompareFloat)

	if !util.IsSorted(f1, util.CompareFloat) || !util.IsPerm(f, f1) {
		t.Error("Bubblesort falhou em ordenar vetor de numeros reais")
	}
}

func TestMergeSort(t *testing.T) {
	x := []util.T{uint32(20), uint32(40), uint32(50), uint32(30), uint32(10)}
	x1 := x
	v := []util.T{"lucas", "joao victor", "joao pedro", "alequi", "maria eduarda"}
	v1 := v
	f := []util.T{float32(2.72), float32(4.57), float32(5.98), float32(3.14), float32(1.18)}
	f1 := f

	sort.Mergesort(x1, 0, len(x1)-1, util.CompareInt)

	if !util.IsSorted(x1, util.CompareInt) || !util.IsPerm(x, x1) {
		t.Error("Mergesort falhou em ordenar vetor de inteiros")
	}

	sort.Mergesort(v1, 0, len(v1)-1, util.CompareString)

	if !util.IsSorted(v1, util.CompareString) || !util.IsPerm(v, v1) {
		t.Error("Mergesort falhou em ordenar vetor de strings")
	}

	sort.Mergesort(f1, 0, len(f1)-1, util.CompareFloat)

	if !util.IsSorted(f1, util.CompareFloat) || !util.IsPerm(f, f1) {
		t.Error("Mergesort falhou em ordenar vetor de numeros reais")
	}
}

func TestQuickSort(t *testing.T) {
	x := []util.T{uint32(20), uint32(40), uint32(50), uint32(30), uint32(10)}
	x1 := x
	v := []util.T{"lucas", "joao victor", "joao pedro", "alequi", "maria eduarda"}
	v1 := v
	f := []util.T{float32(2.72), float32(4.57), float32(5.98), float32(3.14), float32(1.18)}
	f1 := f

	sort.Quicksort(x1, 0, len(x1)-1, util.CompareInt)

	if !util.IsSorted(x1, util.CompareInt) || !util.IsPerm(x, x1) {
		t.Error("Quicksort falhou em ordenar vetor de inteiros")
	}

	sort.Quicksort(v1, 0, len(v1)-1, util.CompareString)

	if !util.IsSorted(v1, util.CompareString) || !util.IsPerm(v, v1) {
		t.Error("Quicksort falhou em ordenar vetor de strings")
	}

	sort.Quicksort(f1, 0, len(f)-1, util.CompareFloat)

	if !util.IsSorted(f1, util.CompareFloat) || !util.IsPerm(f, f1) {
		t.Error("Quicksort falhou em ordenar vetor de numeros reais")
	}
}

func TestMergeSortP(t *testing.T) {
	x := []util.T{uint32(20), uint32(40), uint32(50), uint32(30), uint32(10)}
	x1 := x
	v := []util.T{"lucas", "joao victor", "joao pedro", "alequi", "maria eduarda"}
	v1 := v
	f := []util.T{float32(2.72), float32(4.57), float32(5.98), float32(3.14), float32(1.18)}
	f1 := f

	sort.Mergesort_P(x1, util.CompareInt)

	if !util.IsSorted(x1, util.CompareInt) || !util.IsPerm(x, x1) {
		t.Error("Mergesort Paralelo falhou em ordenar vetor de inteiros")
	}

	sort.Mergesort_P(v1, util.CompareString)

	if !util.IsSorted(v1, util.CompareString) || !util.IsPerm(v, v1) {
		t.Error("Mergesort Paralelo falhou em ordenar vetor de strings")
	}

	sort.Mergesort_P(f1, util.CompareFloat)

	if !util.IsSorted(f1, util.CompareFloat) || !util.IsPerm(f, f1) {
		t.Error("Mergesort Paralelo falhou em ordenar vetor de numeros reais")
	}
}

func TestQuickSortP(t *testing.T) {
	x := []util.T{uint32(20), uint32(40), uint32(50), uint32(30), uint32(10)}
	x1 := x
	v := []util.T{"lucas", "joao victor", "joao pedro", "alequi", "maria eduarda"}
	v1 := v
	f := []util.T{float32(2.72), float32(4.57), float32(5.98), float32(3.14), float32(1.18)}
	f1 := f

	sort.Quicksort_P(x1, 0, len(x1)-1, util.CompareInt)

	if !util.IsSorted(x1, util.CompareInt) || !util.IsPerm(x, x1) {
		t.Error("Quicksort Paralelo falhou em ordenar vetor de inteiros")
	}

	sort.Quicksort_P(v1, 0, len(v1)-1, util.CompareString)

	if !util.IsSorted(v1, util.CompareString) || !util.IsPerm(v, v1) {
		t.Error("Quicksort Paralelo falhou em ordenar vetor de strings")
	}

	sort.Quicksort_P(f1, 0, len(f1)-1, util.CompareFloat)

	if !util.IsSorted(f1, util.CompareFloat) || !util.IsPerm(f, f1) {
		t.Error("Quicksort Paralelo falhou em ordenar vetor de numeros reais")
	}
}

func TestMergeFiles(t *testing.T) {
	util.GenerateFiles("TestMergeFiles.bin", 200, 42)

	file1, err := os.Open("TestMergeFiles.bin")
	if err != nil {
		t.Error("Erro ao abrir o arquivo TestMergeFiles.bin")
	}

	x := util.ReadIntegers(file1, 200)

	sort.Merge_Files("IntegersGo2.bin", "quick-sort", 4, 10, util.ReadIntegers, util.CompareInt, util.FragmentBin, util.WriteIntegers)

	file2, err := os.Open("Sorted.bin")
}
