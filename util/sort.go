package util

type T interface{}


func IsSorted(v []T, cmp func(T, T) bool) bool {
	size := len(v)
	for i := 0; i < size-1; i++ {
		if !cmp(v[i], v[i+1]) && v[i]!=v[i+1] {
			return false
		}
	}
	return true
}


