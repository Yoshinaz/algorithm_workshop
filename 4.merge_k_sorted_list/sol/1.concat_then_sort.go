package sol

import "sort"

func ConcatArray(a [][]int)[]int{
	result := make([]int,0)
	for _,v := range a{
		result = append(result,v...)
	}
	sort.Ints(result)

	return result
}
