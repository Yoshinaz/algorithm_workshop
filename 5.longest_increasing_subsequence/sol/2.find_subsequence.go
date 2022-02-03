package sol

func LISSubsequence(a []int)[]int{
	n:=len(a)
	mic := make([]int,n)
	subIdx := make([]int,n)
	for i:=0;i<n;i++{
		subIdx[i] = i
		for j:=0;j<i;j++{
			if a[i]>a[j] && mic[i]<mic[j]{
				mic[i]=mic[j]
				subIdx[i] = j
			}
		}
		mic[i]++
	}
	result := 0
	ans := make([]int,0)
	idx := 0
	for i,v := range mic{
		if result<v{
			result = v
			idx = i
		}
	}
	for idx != subIdx[idx]{
		ans = append(ans, a[idx])
		idx = subIdx[idx]
	}
	ans = append(ans, a[idx])


	return rev(ans)
}

func rev(s []int)[]int{
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
