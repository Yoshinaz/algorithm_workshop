package sol

import "sort"

func Merge(a [][]int)[]int{
	result := make([]int,0)
	for _,v := range a{
		result = append(result,v...)
	}
	sort.Ints(result)

	return result
}

func merge(a [][]int,l,r int)[]int{
	if l==r{
		return a[l]
	}
	if l+1 == r{
		mergeArray(a[l],a[r])
	}
	m:= (l+r)/2
	left := merge(a,l,m)
	right := merge(a,m+1,r)
	return mergeArray(left,right)
}

func mergeArray(a,b []int)[]int{
	result := make([]int,0)
	i,j := 0,0
	na,nb :=len(a),len(b)
	for ;i<na && j<nb;{
		if a[i]<b[j]{
			result = append(result,a[i])
			i++
		}else{
			result = append(result,b[j])
			j++
		}
	}
	for i<na{
		result = append(result,a[i])
		i++
	}
	for j<nb{
		result = append(result,b[j])
		j++
	}
	return result
}
