package util

func CompareString(a, b T) bool {
	return a.(string) < b.(string)
}

func CompareInt(a, b T) bool {
	return a.(uint32) < b.(uint32)
}

func CompareFloat(a, b T) bool {
	return a.(int) < b.(int) //deixa baixo
}
