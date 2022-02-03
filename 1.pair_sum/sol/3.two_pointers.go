package sol

import "sort"

func TwoPointers2Sum(a []int,k int) bool{
	sort.Ints(a)
	return isSumPairWith2Pointers(a,k)
}

func isSumPairWith2Pointers(a []int,k int) bool {
	l,r := 0,len(a)-1
	for l<r{
		if a[l]+a[r]==k{
			return true
		}
		if a[l]+a[r]<k{
			l++
		}else{
			r--
		}
	}
	return false
}

