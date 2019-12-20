package validbinarytree98

import (
	"container/list"
	"math"

	tree "github.com/go_problem_solving/leetcode/Tree"
	stack "github.com/go_problem_solving/leetcode/datastructures/Stack"
)

// Recursive
func isValidBST_Recursive(root *tree.TreeNode) bool {
	// return isValidBST_Helper(root, nil, nil)
	return isValidBST_Helper(root, math.MinInt64, math.MaxInt64)
}

func isValidBST_Helper(root *tree.TreeNode, min, max int64) bool {
	if root == nil {
		return true
	}
	if root.Val <= int(min) || root.Val >= int(max) {
		return false
	}
	return isValidBST_Helper(root.Left, min, int64(root.Val)) && isValidBST_Helper(root.Right, int64(root.Val), max)
}

// Iterative
func isValidBST_Iterative(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}

	upperLmts := stack.Stack{}
	upperLmts.Push(nil)

	lowerLmts := stack.Stack{}
	lowerLmts.Push(nil)

	nodeStack := stack.Stack{}
	nodeStack.Push(root)

	for !nodeStack.IsEmpty() {
		poped, _ := nodeStack.Pop().(*tree.TreeNode)
		lower, _ := lowerLmts.Pop().(*int)
		upper, _ := upperLmts.Pop().(*int)

		if (lower != nil && poped.Val <= *lower) ||
			(upper != nil && poped.Val >= *upper) {
			return false
		}
		if poped.Right != nil {
			nodeStack.Push(poped.Right)
			lowerLmts.Push(&poped.Val)
			upperLmts.Push(upper)
		}
		if poped.Left != nil {
			nodeStack.Push(poped.Left)
			lowerLmts.Push(lower)
			upperLmts.Push(&poped.Val)
		}
	}
	return true
}

func isValidBST_InOrder(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}

	s := list.New()
	curr := root
	inOrderVal := math.MinInt64

	for s.Len() > 0 || curr != nil {
		if curr != nil {
			s.PushBack(curr)
			curr = curr.Left
		} else {
			poped := s.Remove(s.Back()).(*tree.TreeNode)
			if poped.Val <= int(inOrderVal) {
				return false
			}
			inOrderVal = int(poped.Val)
			curr = poped.Right
		}
	}
	return true
}
