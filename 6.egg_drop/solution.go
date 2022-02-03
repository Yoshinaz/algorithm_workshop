package main

import (
	"eggdrop/sol"
	"fmt"
)

func main(){
	fmt.Println(solution(10,2))
}

func solution(n,k int)int{
	return sol.EggdropMMoveKEgg(n,k)
}
