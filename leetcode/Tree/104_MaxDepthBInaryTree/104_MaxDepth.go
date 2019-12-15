package maxdepth104

import "math"

type maxDepth struct {
	max int
}

//DFS - PreOrder
func maxDepth_Pre(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepth := &maxDepth{
		max: math.MinInt64,
	}
	maxDepth_Pre_Helper(root, 1, maxDepth)
	return maxDepth.max
}

func maxDepth_Pre_Helper(root *TreeNode, depth int, maxDepth *maxDepth) {
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
func maxDepth_Post(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth_Post(root.Left)
	right := maxDepth_Post(root.Right)
	return int(math.Max(float64(left), float64(right)) + 1)
}

/* func inOrder(root *TreeNode, depth int, maxDepth *result) {
	if root.Left == nil && root.Right == nil {
		maxDepth.maxDepth = Max(maxDepth.maxDepth, depth)
		return
	}
	if root.Left != nil {
		inOrder(root.Left, depth+1, maxDepth)
	}
	//visit node
	if root.Right != nil {
		inOrder(root.Right, depth+1, maxDepth)
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
} */
