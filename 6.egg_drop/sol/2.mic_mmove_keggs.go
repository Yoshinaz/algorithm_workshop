package sol

import "fmt"

func EggdropMMoveKEgg(n,k int)int{
	return eggdrop2(n,k)
}

func eggdrop(n,k int)int {
	mic := make([][]int,n+1)
	for i:=0;i<=n;i++ {
		mic[i] = make([]int, k+1)
	}
	m:=0
	for mic[m][k]<n{
		m++
		for i:=1;i<=k;i++{
			mic[m][i] = mic[m-1][i-1]+mic[m-1][i]+1
		}
		fmt.Println(mic[m])
	}
	return m
}


func eggdrop2(n,k int)int {
	mic := make([]int,k+1)
	m:=0
	for mic[k]<n{
		m++
		for i:=k;i>0;i--{
			mic[i] = mic[i-1]+mic[i]+1
		}
	}
	return m
}

