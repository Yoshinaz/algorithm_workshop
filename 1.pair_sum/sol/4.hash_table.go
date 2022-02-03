package sol

func Hash2Sum(a []int,k int) bool{
	seenNums := map[int]bool{}
	for _, num := range a {
		potentialMatch := k - num
		if _, found := seenNums[potentialMatch]; found {
			return true
		}
		seenNums[num] = true
	}
	return false
}

