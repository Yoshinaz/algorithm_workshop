package sol

const MAX_INT = int(^uint(0)>>1)

func CompareMin(a [][]int)[]int{
	idx := make([]int,len(a))
	result := make([]int,0)
	for min:=getMin(a,idx);min!=MAX_INT;min=getMin(a,idx){
		result = append(result,min)
	}

	return result
}

func getMin(a [][]int,idx []int) int{
	min := MAX_INT
	minIdx := 0
	for i,v := range a{
		if idx[i]<len(v) && min>v[idx[i]]{
			min = v[idx[i]]
			minIdx = i
		}
	}
	idx[minIdx]++
	return min
}

