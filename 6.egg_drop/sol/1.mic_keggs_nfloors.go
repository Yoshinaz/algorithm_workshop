package sol

func EggdropKEggNFloor(n,k int)int{
	mic := make([][]int,k)
	for i:=0;i<k;i++ {
		mic[i] = make([]int, n+1)
	}
	for i:=0;i<=n;i++{
		mic[0][i] = i
	}

	for i:=1;i<k;i++{
		for j:=1;j<=n;j++{
			mic[i][j] = n
			for k:=1;k<=j;k++{
				tmp := max(mic[i-1][k-1],mic[i][j-k])+1
				if tmp<mic[i][j]{
					mic[i][j] = tmp
				}
			}
		}
	}

	return mic[k-1][n]
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
