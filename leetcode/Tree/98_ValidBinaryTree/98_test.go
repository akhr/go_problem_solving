package validbinarytree98

import (
	"testing"

	tree "github.com/go_problem_solving/leetcode/Tree"
)

func Test_isValidBST(t *testing.T) {
	tests := []struct {
		items []int
		want  bool
	}{
		{
			items: []int{2, 1, 3},
			want:  true,
		},
		{
			items: []int{5, 1, 4, -1, -1, 3, 6},
			want:  false,
		},
	}
	for _, tt := range tests {
		root := tree.ConstructIntValTree(tt.items, 0)
		t.Run("", func(t *testing.T) {
			if got := isValidBST_Recursive(root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBST_Iterative(t *testing.T) {
	tests := []struct {
		items []int
		want  bool
	}{
		{
			items: []int{2, 1, 3},
			want:  true,
		},
		{
			items: []int{5, 1, 4, -1, -1, 3, 6},
			want:  false,
		},
	}
	for _, tt := range tests {
		root := tree.ConstructIntValTree(tt.items, 0)
		t.Run("", func(t *testing.T) {
			if got := isValidBST_Iterative(root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBST_InOrder(t *testing.T) {
	tests := []struct {
		items []int
		want  bool
	}{
		{
			items: []int{2, 1, 3},
			want:  true,
		},
		{
			items: []int{5, 1, 4, -1, -1, 3, 6},
			want:  false,
		},
	}
	for _, tt := range tests {
		root := tree.ConstructIntValTree(tt.items, 0)
		t.Run("", func(t *testing.T) {
			if got := isValidBST_InOrder(root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
