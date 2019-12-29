package sort

import "fmt"

func quickSort(a []int) {
	qHelper(a, 0, len(a)-1)
}

func qHelper(a []int, s, e int) {
	if s >= e {
		return
	}
	p := partition(a, s, e)
	qHelper(a, s, p-1)
	qHelper(a, p+1, e)
}

func partition(a []int, start, end int) int {
	pivot := a[end]
	p := start
	for i := start; i < end; i++ {
		if a[i] < pivot {
			a[i], a[p] = a[p], a[i] //swap smaller elements to front of pivot
			p++
		}
	}
	a[p], a[end] = a[end], a[p] // swap pivot to its final partitionIndex
	fmt.Println(a)
	return p
}
