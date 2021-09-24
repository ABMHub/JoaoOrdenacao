package util

func CompareString(a, b T) bool {
	return a.(string) < b.(string)
}

func CompareInt(a, b T) bool {
	return a.(int) < b.(int)
}

func CompareFloat(a, b T) bool {
	return a.(int) < b.(int)
}
