package main

import (
	"fmt"
	"ksort/sol"
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
	return sol.ConcatArray(a)
}



