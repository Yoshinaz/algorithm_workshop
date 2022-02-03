package main


func countPrimes(n int) int{
	if n<2{
		return 0
	}
	nonPrime := make([]bool,n+1)
	numOfNonPrimes := 1
	for i:=2;i<=n;i++{
		if nonPrime[i]{
			continue
		}
		for j:=i*2;j<=n;j+=i{
			if !nonPrime[j]{
				nonPrime[j] = true
				numOfNonPrimes++
			}
		}
	}
	return n - numOfNonPrimes
}

