package main

import (
	"fmt"
	"samanun/algorithm_workshop/5.longest_increasing_subsequence/sol"
)


func main(){
	input := []int{10, 22, 9, 33, 21, 50, 41, 60, 80}
	fmt.Println(solution_findlength(input))
	fmt.Println(solution_findsubsequence(input))
}

func solution_findlength(a []int)int{
	return sol.LISLength(a)
}

func solution_findsubsequence(a []int)[]int{
	return sol.LISSubsequence(a)
}
