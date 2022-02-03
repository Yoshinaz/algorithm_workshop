package main

func insertionsort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		for j := i;j > 0;j-- {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
		}
	}
}
