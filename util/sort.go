package util

type T interface{}

func IsSorted(v []T, cmp func(T, T) bool) bool {
	size := len(v)
	for i := 0; i < size-1; i++ {
		if !cmp(v[i], v[i+1]) && v[i] != v[i+1] {
			return false
		}
	}
	return true
}

func IsPerm(v1, v2 []T) bool {
	if len(v1) != len(v2) {
		return false
	}

	hash := make(map[T]int)

	for i := 0; i < len(v1); i++ {
		hash[v1[i]]++
	}

	for i := 0; i < len(v2); i++ {
		hash[v2[i]]--

		if hash[v2[i]] == 0 {
			delete(hash, v2[i])
		}
	}

	if len(hash) == 0 {
		return true
	}
	return false
}
