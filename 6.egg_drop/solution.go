package main

import (
	"fmt"
	"samanun/algorithm_workshop/6.egg_drop/sol"
)

func main(){
	fmt.Println(solution(10,2))
}

func solution(n,k int)int{
	return sol.EggdropMMoveKEgg(n,k)
}
