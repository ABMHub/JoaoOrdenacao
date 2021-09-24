package util

import(
	"sync"
)

var __T_counter__ int = 0
var __T_limit int = 0

func SetThreadLimit(limit int){
	if(limit >= 1){
		__T_limit=limit
	}
}

func ResetThreadCounter(){
	__T_counter__=0
}

func Semaforo(groutine func()){
	if(__T_counter__ < __T_limit){
		var tg sync.WaitGroup
		tg.Add(1)

		__T_counter__++
		
		go func(){
			defer tg.Done()
			groutine()
		}()
		
		tg.Wait()	
		
		__T_counter__--
	}else{
		groutine()	
	}
}	