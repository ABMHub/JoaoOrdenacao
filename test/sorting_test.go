package test

import (
	"sortAlgorithms/sort"
	"sortAlgorithms/util"
	"testing"
)

var strings []util.T = []util.T{"lucas", "joao victor", "joao pedro", "aleq", "maria eduarda"}
var integers []util.T = []util.T{20, 5, 31, 42, 14}
var integers_sorted []util.T = []util.T{10, 20, 30, 40, 50}
var integers_reverse_sorted []util.T = []util.T{50, 40, 30, 20, 10}

func TestCompare(t *testing.T) {
	x := integers
	v := strings

	// Testes da CompareInt
	if !util.CompareInt(x[1], x[2]) {
		t.Errorf("Erro no CompareInt, %d < %d eh falso", x[1], x[2])
	}

	if !util.CompareInt(x[2], x[3]) {
		t.Errorf("Erro no CompareInt, %d < %d eh falso", x[2], x[3])
	}

	if util.CompareInt(x[4], x[4]) {
		t.Errorf("Erro no CompareInt, %d < %d eh verdadeiro", x[4], x[4])
	}

	// Testes da CompareString
	if !util.CompareString(v[3], v[1]) {
		t.Errorf("Erro no CompareString, %s < %s eh falso", v[3], v[1])
	}

	if !util.CompareString(v[2], v[1]) {
		t.Errorf("Erro no CompareString, %s < %s eh falso", v[2], v[1])
	}

	if util.CompareString(v[4], v[4]) {
		t.Errorf("Erro no CompareString, %s < %s eh verdadeiro", v[4], v[4])
	}

	// Testando a funcao IsSorted
	if !util.IsSorted(integers_sorted, util.CompareInt) {
		t.Error("Erro, array [10, 20, 30, 40, 50] eh indicado como nao ordenado")
	}

	if util.IsSorted(integers, util.CompareInt) {
		t.Error("Erro, array [20, 5, 31, 42, 14] eh indicado como ordenado")
	}

	if util.IsSorted(integers_reverse_sorted, util.CompareInt) {
		t.Error("Erro, array [50, 40, 30, 20, 10] eh indicado como ordenado")
	}
}

func TestBubbleSort(t *testing.T) {
	var x []util.T
	copy(x, integers_sorted)
	sort.Bubblesort(x, util.CompareInt)
	if !util.IsSorted(x, util.CompareInt) {
		t.Error("Bubblesort falhou em ordenar")
	}

	copy(x, integers_reverse_sorted)
	sort.Bubblesort(x, util.CompareInt)
	if !util.IsSorted(x, util.CompareInt) {
		t.Error("Bubblesort falhou em ordenar")
	}

	copy(x, integers)
	sort.Bubblesort(x, util.CompareInt)
	if !util.IsSorted(x, util.CompareInt) {
		t.Error("Bubblesort falhou em ordenar")
	}

	copy(x, strings)
	sort.Bubblesort(x, util.CompareString)
	if !util.IsSorted(x, util.CompareString) {
		t.Error("Bubblesort falhou em ordenar")
	}
}

func TestMergeSort(t *testing.T) {
	var x []util.T
	copy(x, integers_sorted)
	sort.Mergesort(x, 0, len(x)-1, util.CompareInt)
	if !util.IsSorted(x, util.CompareInt) {
		t.Error("Merge falhou em ordenar")
	}

	copy(x, integers_reverse_sorted)
	sort.Mergesort(x, 0, len(x)-1, util.CompareInt)
	if !util.IsSorted(x, util.CompareInt) {
		t.Error("Merge falhou em ordenar")
	}

	copy(x, integers)
	sort.Mergesort(x, 0, len(x)-1, util.CompareInt)
	if !util.IsSorted(x, util.CompareInt) {
		t.Error("Merge falhou em ordenar")
	}

	copy(x, strings)
	sort.Mergesort(x, 0, len(x)-1, util.CompareString)
	if !util.IsSorted(x, util.CompareString) {
		t.Error("Merge falhou em ordenar")
	}
}

func TestQuickSort(t *testing.T) {
	var x []util.T
	copy(x, integers_sorted)
	sort.Quicksort(x, 0, len(x)-1, util.CompareInt)
	if !util.IsSorted(x, util.CompareInt) {
		t.Error("Quick falhou em ordenar")
	}

	copy(x, integers_reverse_sorted)
	sort.Quicksort(x, 0, len(x)-1, util.CompareInt)
	if !util.IsSorted(x, util.CompareInt) {
		t.Error("Quick falhou em ordenar")
	}

	copy(x, integers)
	sort.Quicksort(x, 0, len(x)-1, util.CompareInt)
	if !util.IsSorted(x, util.CompareInt) {
		t.Error("Quick falhou em ordenar")
	}

	copy(x, strings)
	sort.Quicksort(x, 0, len(x)-1, util.CompareString)
	if !util.IsSorted(x, util.CompareString) {
		t.Error("Bubblesort falhou em ordenar")
	}
}
