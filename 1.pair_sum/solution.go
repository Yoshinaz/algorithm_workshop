package main

import (
	"fmt"
	"pairsum/sol"
)

func main(){
	fmt.Println(solution([]int{1,2,3,4},7))
}

func solution(a []int,k int) bool{
	return sol.Hash2Sum(a,k)
}

