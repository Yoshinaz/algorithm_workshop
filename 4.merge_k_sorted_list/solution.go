package main

import (
	"fmt"
	"samanun/algorithm_workshop/4.merge_k_sorted_list/sol"
)

func main(){
	input := [][]int{
		{
			1,4,6,8,
		},
		{
			2,5,100,101,
		},
		{
			2,6,99,102,
		},
	}

	fmt.Println(solution(input))
}

func solution(a [][]int)[]int{
	return sol.CompareMinHeap(a)
}



