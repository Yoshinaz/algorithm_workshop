package sol

import "sort"

func Solution(a []int,k int) bool{
	sort.Ints(a)
	for i,v := range a{
		b := remove(a,i)
		if twoPointers2Sum(b,k-v){
			return true
		}
	}
	return false
}

func remove(slice []int, s int) []int {
	b := make([]int, len(slice))
	copy(b, slice)
	return append(b[:s], b[s+1:]...)
}


func twoPointers2Sum(a []int,k int) bool{
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
