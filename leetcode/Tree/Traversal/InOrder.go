package traversal

import (
	"fmt"

	stack "github.com/go_problem_solving/leetcode/Stack"
)

//Recursive
func InOrder_Recusive(root *TreeNode) {
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

func InOrder_Iterative_1_Stack(root *TreeNode) {
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
			node, _ := poped.(*TreeNode)
			fmt.Println(node.Val) //Visit node
			curr = node.Right
		}
	}
}
