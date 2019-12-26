package sort

func mergeSort(a []int) {
	mergeSortHelper(a, 0, len(a)-1)
}

func mergeSortHelper(a []int, start, end int) {
	if start >= end {
		return
	}
	mid := start + (end-start)/2
	mergeSortHelper(a, start, mid)
	mergeSortHelper(a, mid+1, end)
	merge(a, start, mid, end)
}

func merge(a []int, start, mid, end int) {
	hArr := []int{}
	l := start
	r := mid + 1
	for l <= mid && r <= end {
		if a[l] < a[r] {
			hArr = append(hArr, a[l])
			l++
		} else {
			hArr = append(hArr, a[r])
			r++
		}
	}
	for l <= mid {
		hArr = append(hArr, a[l])
		l++
	}
	for r <= end {
		hArr = append(hArr, a[r])
		r++
	}

	// copy helperArr contents to original array
	copy(a[start:], hArr[:])
}
