package sort

import(
	"sortAlgorithms/util"
)

func Mergesort_P(arr []util.T, cmp func(util.T, util.T) bool){
	
	p_threads := util.GetThreadLimit() + 1
	var fin, part, npart, arrSize int
	bF := make([]util.Pair, p_threads + 1)
	bolH := make([]bool, p_threads + 1)
	arrSize = len(arr)
	fin = arrSize - 1
	part, npart = (fin + 1)/p_threads, (fin+1)%p_threads

	if(npart==fin){
		npart--
	}

	bF[0] = util.Pair{0,part+npart-1}

	for j := 1;; j++{
		bolH[j-1]=false

		go func(idx int){
			defer func(){
				bolH[idx]=true
			}()
			Mergesort(arr, bF[idx].Fst.(int), bF[idx].Snd.(int), cmp)
		}(j-1)
		if(bF[j-1].Snd.(int) == fin){
			break
		}
		bF[j] = util.Pair{bF[j-1].Snd.(int) + 1, bF[j-1].Snd.(int) + part}
	}

	for bF[0].Snd.(int) - bF[0].Fst.(int) + 1 < len(arr){
		for i := 0; i < p_threads; i++{
			if bolH[i]{
				for j := i + 1; j < p_threads; j++{
					if(bF[i].Snd.(int) == bF[j].Fst.(int)-1 && bolH[j]){
						bolH[i]=false
						bolH[j]=false
						go func(idxa, idxb int){
							defer func(){
								bF[idxa].Snd=bF[idxb].Snd
								bolH[idxa]=true
							}()
							merge(arr, bF[idxa].Fst.(int), bF[idxa].Snd.(int), 
														bF[idxb].Snd.(int), cmp)
						}(i,j)
						break
					}
				}
			}
		}
	}
}