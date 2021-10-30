package util

import(
	"sync"
)

var __T_counter__ int = 0
var __T_limit int = 0

func SetThreadLimit(limit int){
	if(limit >= 0){
		__T_limit=limit
	}
}

func GetThreadLimit()int{
	return __T_limit
}

func ResetThreadCounter(){
	__T_counter__=0
}

func Semaforo(n int, groutine ...func()){
	if(__T_counter__ < __T_limit){
		var tg sync.WaitGroup
		tg.Add(1)

		__T_counter__++
		
		go func(){
			defer tg.Done()
			for i := 0; i < n; i++{
				groutine[i]()
			}
		}()
		for i := n; i < len(groutine); i++{
			groutine[i]()
		}
		tg.Wait()	
		
		__T_counter__--
	}else{
		for i := 0; i < len(groutine); i++{
			groutine[i]()
		}
	}
}	