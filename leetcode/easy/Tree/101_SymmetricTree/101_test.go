package symmetricTree101

import (
	"testing"

	tree "github.com/go_problem_solving/leetcode/Tree"
)

func Test_isSymmetric(t *testing.T) {
	tests := []struct {
		items []int
		want  bool
	}{
		{
			items: []int{1, 2, 2, 3, 4, 4, 3},
			want:  true,
		},
		{
			items: []int{1, 2, 2, 5, 4, 4, 3},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			root := tree.ConstructIntValTree(tt.items, 0)
			if got := isSymmetric(root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSymmetric_2(t *testing.T) {
	tests := []struct {
		items []int
		want  bool
	}{
		{
			items: []int{1, 2, 2, 3, 4, 4, 3},
			want:  true,
		},
		{
			items: []int{1, 2, 2, 5, 4, 4, 3},
			want:  false,
		},
		{
			items: []int{9, -42, -42, -1, 76, 76, -1, -1, 13, -1, 13},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			root := tree.ConstructIntValTree(tt.items, 0)
			if got := isSymmetric_2(root); got != tt.want {
				t.Errorf("isSymmetric_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
