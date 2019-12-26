package sort

import "fmt"

// i represents start of un-sorted partion
// j represents start of sorted portion and ends before i
// At each iteration of `i` we need to find the right element for it and add `i` to sorted portion
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
			fmt.Printf("i : %d, j : %d, arr : %v\n", i, j, arr)
		}
	}
}

// [3 || 5 6 2 4 1]
// [3 5 || 6 2 4 1]
// [3 5 6 || 2 4 1]
// [2 5 6 || 3 4 1]
// [2 3 6 || 5 4 1]
// [2 3 5 || 6 4 1]
// [2 3 5 6 || 4 1]
// [2 3 5 6 || 4 1]
// [2 3 4 6 5 || 1]
// [1 3 4 5 6 || 2]
// [1 2 4 5 6 || 3]
// [1 2 3 5 6 || 4]
// [1 2 3 4 6 || 5]
// [1 2 3 4 5 6 ||]

// i : 1, j : 0, arr : [|| 3 5 6 2 4 1]
// i : 2, j : 0, arr : [3 || 5 6 2 4 1]
// i : 2, j : 1, arr : [3 || 5 6 2 4 1]
// i : 3, j : 0, arr : [2 5 6 || 3 4 1]
// i : 3, j : 1, arr : [2 3 6 || 5 4 1]
// i : 3, j : 2, arr : [2 3 5 || 6 4 1]
// i : 4, j : 0, arr : [2 3 5 6 || 4 1]
// i : 4, j : 1, arr : [2 3 5 6 || 4 1]
// i : 4, j : 2, arr : [2 3 4 6 || 5 1]
// i : 4, j : 3, arr : [2 3 4 5 || 6 1]
// i : 5, j : 0, arr : [1 3 4 5 6 || 2]
// i : 5, j : 1, arr : [1 2 4 5 6 || 3]
// i : 5, j : 2, arr : [1 2 3 5 6 || 4]
// i : 5, j : 3, arr : [1 2 3 4 6 || 5]
// i : 5, j : 4, arr : [1 2 3 4 5 || 6]
