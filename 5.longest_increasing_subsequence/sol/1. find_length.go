package sol

func LISLength(a []int)int{
	n:=len(a)
	mic := make([]int,n)
	for i:=0;i<n;i++{
		for j:=0;j<i;j++{
			if a[i]>a[j] && mic[i]<mic[j]{
				mic[i]=mic[j]
			}
		}
		mic[i]++
	}
	result := max(mic)
	return result
}

func max(a []int)int{
	result := 0
	for _,v := range a{
		if result<v{
			result = v
		}
	}
	return result
}
