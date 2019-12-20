package sortedarraytobst

import (
	"testing"

	tree "github.com/go_problem_solving/leetcode/Tree"
	traversal "github.com/go_problem_solving/leetcode/Tree/Traversal"
)

func Test_sortedArrayToBST(t *testing.T) {
	tests := []struct {
		nums []int
		want *tree.TreeNode
	}{
		{
			nums: []int{-10, -3, 0, 5, 9},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := sortedArrayToBST(tt.nums)
			t.Log("Constructed Tree :: ")
			traversal.InOrder_Recusive(got)
			t.Logf("Input Nums :: %+v", tt.nums)
		})
	}
}
