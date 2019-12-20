package traversal

import (
	"testing"

	tree "github.com/go_problem_solving/leetcode/Tree"
)

func Test_PostOrder_Recursive(t *testing.T) {
	tests := []struct {
		items []int
	}{
		{
			items: []int{3, 9, 20, -1, -1, 15, 7},
		},
		{
			items: []int{},
		},
		{
			items: []int{1, 2},
		},
		{
			items: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		root := tree.ConstructIntValTree(tt.items, 0)
		t.Run("", func(t *testing.T) {
			PostOrder_Recusive(root)
		})
	}
}

/* func Test_PostOrder_Iterative_1_Stack(t *testing.T) {
	tests := []struct {
		items []int
	}{
		{
			items: []int{3, 9, 20, -1, -1, 15, 7},
		},
		{
			items: []int{},
		},
		{
			items: []int{1, 2},
		},
		{
			items: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		root := constructTree(tt.items, 0)
		t.Run("", func(t *testing.T) {
			PostOrder_Iterative_1_Stack(root)
		})
	}
} */

func Test_PostOrder_Iterative_2_Stack(t *testing.T) {
	tests := []struct {
		items []int
	}{
		{
			items: []int{3, 9, 20, -1, -1, 15, 7},
		},
		{
			items: []int{},
		},
		{
			items: []int{1, 2},
		},
		{
			items: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		root := tree.ConstructIntValTree(tt.items, 0)
		t.Run("", func(t *testing.T) {
			PostOrder_Iterative_2_Stack(root)
		})
	}
}
