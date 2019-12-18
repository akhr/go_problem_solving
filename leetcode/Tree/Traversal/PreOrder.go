package traversal

import (
	"fmt"

	stack "github.com/go_problem_solving/leetcode/Stack"
)

//Recursive
func PreOrder_Recusive(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val) //Visit node
	if root.Left != nil {
		PreOrder_Recusive(root.Left)
	}
	if root.Right != nil {
		PreOrder_Recusive(root.Right)
	}
}

func PreOrder_Iterative_1_Stack(root *TreeNode) {
	if root == nil {
		return
	}
	s := stack.Stack{}
	s.Push(root)
	for !s.IsEmpty() {
		poped := s.Pop()
		if node, ok := poped.(*TreeNode); ok {
			fmt.Println(node.Val) //Visit node
			if node.Right != nil {
				s.Push(node.Right)
			}
			if node.Left != nil {
				s.Push(node.Left)
			}
		}
	}
}