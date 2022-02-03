package main

func fibo(n int) int{
	a := make([]int,n+1)
	a[0],a[1]=1,1
	for i:=2;i<=n;i++{
		a[i]=a[i-1]+a[i-2]
	}
	return a[n]
}


func fibo2(n int) int {
	if n<2{
		return 1
	}
	return fibo2(n-1)+fibo2(n-2)
}
