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

type Reddit struct {
    Title     string  `json:"title"`
    Author    string  `json:"author"`
    Timestamp float64 `json:"timestamp"`
    ID        string  `json:"id"`
    Likes     int     `json:"likes"`
}
