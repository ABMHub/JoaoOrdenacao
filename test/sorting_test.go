package test

import (
	"sortAlgorithms/sort"
	"sortAlgorithms/util"
	"testing"
)

var v [5]string = [5]string{"lucas", "joao victor", "joao pedro", "aleq", "maria eduarda"}

func TestCompare(t *testing.T) {
	x := []util.T{10, 20, 30, 40, 50}

	if util.CompareInt(x[0], x[1]) {
		t.Errorf("Erro no CompareInt, %d < %d eh falso", x[0], x[1])
	}

	if util.CompareInt(x[2], x[3]) {
		t.Errorf("Erro no CompareInt, %d < %d eh falso", x[2], x[3])
	}

	if util.CompareInt(x[4], x[4]) {
		t.Errorf("Erro no CompareInt, %d < %d eh verdadeiro", x[4], x[4])
	}

	if util.CompareString(v[3], v[1]) {
		t.Errorf("Erro no CompareString, %s < %s eh falso", v[3], v[1])
	}

	if util.CompareString(v[2], v[1]) {
		t.Errorf("Erro no CompareString, %s < %s eh falso", v[2], v[1])
	}

	if util.CompareString(v[4], v[4]) {
		t.Errorf("Erro no CompareString, %s < %s eh verdadeiro", v[4], v[4])
	}
}

func TestSorted(t *testing.T) {
	x := []util.T{10, 20, 30, 40, 50}
	if !util.IsSorted(x, util.CompareInt) {
		t.Error("Erro, array [10, 20, 30, 40, 50] eh indicado como nao ordenado")
	}

	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}

	sort.Bubblesort(x, util.CompareInt)

	if !util.IsSorted(x, util.CompareInt) {
		t.Error("Bubblesort falhou em ordenar")
	}

	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}

	sort.Mergesort(x, 0, len(x)-1)

	if !util.IsSorted(x, util.CompareInt) {
		t.Error("Merge falhou em ordenar")
	}
}
