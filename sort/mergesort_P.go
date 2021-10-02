package sort

import(
	"sortAlgorithms/util"
)

func Mergesort_P(arr []util.T,	 cmp func(util.T, util.T) bool){
	
	p_threads := util.GetThreadLimit() + 1
	var part, npart int // ultimo idx do arr, particoes, arruma as particoes, ...
	bF := make([]util.Pair, p_threads + 1) // guarda begin e end de cada particao.
	bolH := make([]bool, p_threads + 1) // fala se a particao esta ou nao preparada para modificacao
	
	part, npart = (fin-begin+1)/p_threads, (fin-begin+1)%p_threads // pega as particoes
	
	bF[0] = util.Pair{0,part+npart-1} // inicializa a primeira particao

	for j := 1;; j++{
		bolH[j-1]=false

		go func(idx int){
			defer func(){
				bolH[idx]=true // no fim da thread, fala q a particao ta preparada.
			}()
			Mergesort(arr, bF[idx].Fst.(int), bF[idx].Snd.(int), cmp)
		}(j-1)
		if(bF[j-1].Snd.(int) == fin){ // checagem de ultima particao.
			break
		}
		bF[j] = util.Pair{bF[j-1].Snd.(int) + 1, bF[j-1].Snd.(int) + part}
		// calculo para proxima particao.
	}

	for bF[0].Snd.(int) - bF[0].Fst.(int) + 1 < fin-begin+1{ 
		for i := 0; i < p_threads; i++{ // checa todas as particoes.
			if bolH[i]{ // particao ta pronta
				for j := i + 1; j < p_threads; j++{ // checa pela proxima particao pronta
					if(bF[i].Snd.(int) == bF[j].Fst.(int)-1 && bolH[j]){ // particoes "tocando" + particao pronta.
						bolH[i]=false
						bolH[j]=false
						go func(idxa, idxb int){
							defer func(){
								bF[idxa].Snd=bF[idxb].Snd // particao A "engole" a particao B 
								bolH[idxa]=true // so a primeira particao fica pronta.
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