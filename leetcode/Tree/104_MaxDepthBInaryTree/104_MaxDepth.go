package maxdepth104

import (
	"math"

	tree "github.com/go_problem_solving/leetcode/Tree"
)

type maxDepth struct {
	max int
}

//DFS - PreOrder
func maxDepth_Pre(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepth := &maxDepth{
		max: math.MinInt64,
	}
	maxDepth_Pre_Helper(root, 1, maxDepth)
	return maxDepth.max
}

func maxDepth_Pre_Helper(root *tree.TreeNode, depth int, maxDepth *maxDepth) {
	if root.Left == nil && root.Right == nil {
		maxDepth.max = int(math.Max(float64(maxDepth.max), float64(depth)))
	}
	if root.Left != nil {
		maxDepth_Pre_Helper(root.Left, depth+1, maxDepth)
	}
	if root.Right != nil {
		maxDepth_Pre_Helper(root.Right, depth+1, maxDepth)
	}
}

// DFS - PostOrder
func maxDepth_Post(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth_Post(root.Left)
	right := maxDepth_Post(root.Right)
	return int(math.Max(float64(left), float64(right)) + 1)
}
