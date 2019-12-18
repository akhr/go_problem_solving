package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ConstructIntValTree(items []int, index int) *TreeNode {
	if index > len(items)-1 {
		return nil
	}

	if items[index] == -1 {
		return nil
	}

	node := &TreeNode{Val: items[index]}
	node.Left = ConstructIntValTree(items, getLeftChildIndx(index))
	node.Right = ConstructIntValTree(items, getRightChildIndx(index))
	return node
}

func getParentIndx(childIndx int) int {
	return (childIndx - 1) / 2
}

func getLeftChildIndx(parentIndx int) int {
	return (parentIndx * 2) + 1
}

func getRightChildIndx(parentIndx int) int {
	return (parentIndx * 2) + 2
}

func (t *TreeNode) string() string {
	var left, right int
	if t.Left != nil {
		left = t.Left.Val
	}
	if t.Right != nil {
		right = t.Right.Val
	}
	return fmt.Sprintf("node.Val::%d, node.Left.Val::%d, node.Right.Val::%d", t.Val, left, right)
}
