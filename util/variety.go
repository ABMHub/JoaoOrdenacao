package util

func Min(a, b int)int{
	if(a < b){
		return a 
	}
	return b
}

func Max(a, b int)int{
	if(a > b){
		return a 
	}
	return b
}

func Mdc(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Mmc(a, b int) int {
	return (a * b)/Mdc(a, b)
}