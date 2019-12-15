package maxdepth104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructTree(items []int, index int) *TreeNode {
	if index > len(items)-1 {
		return nil
	}

	if items[index] == -1 {
		return nil
	}

	node := &TreeNode{Val: items[index]}
	node.Left = constructTree(items, getLeftChildIndx(index))
	node.Right = constructTree(items, getRightChildIndx(index))
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
