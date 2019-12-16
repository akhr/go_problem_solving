package traversal

import (
	"fmt"

	stack "github.com/go_problem_solving/leetcode/Stack"
)

//Recursive
func PostOrder_Recusive(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		PostOrder_Recusive(root.Left)
	}
	if root.Right != nil {
		PostOrder_Recusive(root.Right)
	}
	fmt.Println(root.Val) //Visit node
}

/* func PostOrder_Iterative_1_Stack(root *TreeNode) {
	if root == nil {
		return
	}
	s := stack.Stack{}
	curr := root

	if curr != nil || !s.IsEmpty() {
		if curr != nil {
			s.Push(curr)
			curr = curr.Left
		} else {
			poped := s.Peek().(*TreeNode)
			if poped.Right != nil {
				curr = poped.Right
			} else {
				fmt.Println(poped.Val) //Visit node
			}
		}
	}
} */

func PostOrder_Iterative_2_Stack(root *TreeNode) {
	if root == nil {
		return
	}
	s1 := stack.Stack{}
	s2 := stack.Stack{}
	s1.Push(root)

	for !s1.IsEmpty() {
		poped := s1.Pop()
		node, _ := poped.(*TreeNode)

		s2.Push(node)
		if node.Left != nil {
			s1.Push(node.Left)
		}
		if node.Right != nil {
			s1.Push(node.Right)
		}
	}

	for !s2.IsEmpty() {
		poped := s2.Pop()
		node, _ := poped.(*TreeNode)
		fmt.Println(node.Val) //Visit node
	}
}
