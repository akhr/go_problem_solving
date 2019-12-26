package sort

import "fmt"

func bubbleSort(arr []int) { //compare adj elements and do it n times. n * n(adj compares) = O(n^2)
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			fmt.Println(arr)
		}
	}
}
