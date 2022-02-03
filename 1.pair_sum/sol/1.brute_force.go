package sol

func BruteForce2Sum(a []int,k int) bool{
	n:=len(a)
	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{
			if a[i]+a[j] == k{
				return true
			}
		}
	}
	return false
}
