package sortedarraytobst

import tree "github.com/go_problem_solving/leetcode/Tree"

func sortedArrayToBST(nums []int) *tree.TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, i, j int) *tree.TreeNode {
	if i > j {
		return nil
	}
	mid := (i + j) / 2
	node := &tree.TreeNode{Val: nums[mid]}
	node.Left = helper(nums, i, mid-1)
	node.Right = helper(nums, mid+1, j)
	return node
}

// From LeetCode
func sortedArrayToBST_2(nums []int) *tree.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	return &tree.TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}
