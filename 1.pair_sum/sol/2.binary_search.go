package sol

import "sort"


func BinarySearch2Sum(a []int,k int) bool {
	sort.Ints(a)
	for i,v := range a{
		b := remove(a,i)
		if binarySearch(b,k-v){
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

func binarySearch(a []int,k int) bool {
	l:=0
	r:=len(a)-1
	for l<=r{
		m := (l+r)/2
		if a[m]==k{
			return true
		}
		if k<a[m]{
			r = m-1
		}
		if k>a[m]{
			l = m+1
		}
	}
	return false
}

func binarySearch2(a []int,k int) bool {
	l:=0
	r:=len(a)-1
	return binarySearchRecur(a,k,l,r)
}

func binarySearchRecur(a []int,k,l,r int) bool{
	m := (l+r)/2
	if a[m]==k{
		return true
	}
	if k<a[m]{
		return binarySearchRecur(a,k,l,m-1)
	}else {
		return binarySearchRecur(a,k,m+1,r)
	}
}
