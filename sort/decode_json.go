package sort

import (
    "encoding/json"
    "os"
    "io/ioutil"
    "sortAlgorithms/util"
)

func SortJson(file *os.File,cmp util.Compare) {
    var arr []util.Reddit
    jsonData, _ := ioutil.ReadAll(file)
    json.Unmarshal(jsonData, &arr)    
    slice := make([]util.T, len(arr))
    for i := range arr {
        slice[i] = arr[i]
    }
    Mergesort_P(slice, cmp)
    res,_ := json.MarshalIndent(slice, "", " ")
    ioutil.WriteFile("Sorted.json", res, os.ModePerm)    
}