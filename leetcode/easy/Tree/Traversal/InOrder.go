package traversal

import (
	"fmt"

	tree "github.com/go_problem_solving/leetcode/Tree"
	stack "github.com/go_problem_solving/leetcode/datastructures/Stack"
)

//Recursive
func InOrder_Recusive(root *tree.TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		InOrder_Recusive(root.Left)
	}
	fmt.Println(root.Val) //Visit node
	if root.Right != nil {
		InOrder_Recusive(root.Right)
	}
}

func InOrder_Iterative_1_Stack(root *tree.TreeNode) {
	if root == nil {
		return
	}
	s := stack.Stack{}
	curr := root
	for curr != nil || !s.IsEmpty() {
		if curr != nil {
			s.Push(curr)
			curr = curr.Left
		} else {
			poped := s.Pop()
			node, _ := poped.(*tree.TreeNode)
			fmt.Println(node.Val) //Visit node
			curr = node.Right
		}
	}
}
