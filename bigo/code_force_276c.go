package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solution(){
	in := bufio.NewReader(os.Stdin)
	var n, q,l,r int
	fmt.Fscan(in, &n,&q)
	a := make([]int,n)
	for i:=0;i<n;i++{
		fmt.Fscan(in,&a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	tmp := make([]int,n+1)
	for i:=0;i<q;i++{
		fmt.Fscan(in,&l,&r)
		tmp[l-1]++
		tmp[r]--
	}

	count := make([]int,n)
	count[0] = tmp[0]
	for i:=1;i<n;i++{
		count[i] = count[i-1]+tmp[i]
	}

	sort.Slice(count, func(i, j int) bool {
		return count[i] > count[j]
	})

	ans := int64(0)
	for i:=0;i<n;i++{
		ans+= int64(count[i])*int64(a[i])
	}
	fmt.Println(ans)
}
