package main

import (
	"fmt"
	"samanun/algorithm_workshop/1-2sum/sol"
)

func main(){
	fmt.Println(solution([]int{1,2,3,4},7))
}

func solution(a []int,k int) bool{
	return sol.Hash2Sum(a,k)
}

