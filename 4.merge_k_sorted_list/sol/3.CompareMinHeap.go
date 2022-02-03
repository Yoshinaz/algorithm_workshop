package sol

import "container/heap"

type data struct{
	value int
	idx int
	index int // The index of the item in the heap.
}

func CompareMinHeap(a [][]int)[]int{
	idx := make([]int,len(a))
	result := make([]int,0)
	var minHeap PriorityQueue
	for i,v := range a{
		heap.Push(&minHeap,&data{
			value: v[0],
			idx: i,
		})
	}
	for minHeap.Len()>0{
		min := (heap.Pop(&minHeap)).(*data)
		result = append(result,min.value)
		idx[min.idx]++
		if idx[min.idx]<len(a[min.idx]){
			heap.Push(&minHeap,&data{
				value: a[min.idx][idx[min.idx]],
				idx: min.idx,
			})
		}
	}

	return result
}



