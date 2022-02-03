package main

import (
	"fmt"
	"lis/sol"
)


func main(){
	input := []int{10, 22, 9, 33, 21, 50, 41, 60, 80}
	fmt.Println(solutionFindLength(input))
	fmt.Println(solutionFindSubsequence(input))
}

func solutionFindLength(a []int)int{
	return sol.LISLength(a)
}

func solutionFindSubsequence(a []int)[]int{
	return sol.LISSubsequence(a)
}
