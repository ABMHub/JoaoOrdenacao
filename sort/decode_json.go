package sort

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sortAlgorithms/util"
)

func SortJson(file *os.File, cmp util.Compare) {
	var arr []util.Reddit
	jsonData, _ := ioutil.ReadAll(file)
	json.Unmarshal(jsonData, &arr)
	slice := make([]util.T, len(arr))
	for i := range arr {
		slice[i] = arr[i]
	}
	util.SetThreadLimit(10)
	Mergesort_P(slice, cmp)
	res, _ := json.MarshalIndent(slice, "", " ")
	ioutil.WriteFile("Sorted.json", res, os.ModePerm)
}
