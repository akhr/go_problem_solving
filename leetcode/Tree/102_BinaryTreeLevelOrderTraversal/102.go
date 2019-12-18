package binarytreeleveltraversal102

import (
	"container/list"

	queue "github.com/go_problem_solving/leetcode/Queue"
	tree "github.com/go_problem_solving/leetcode/Tree"
)

func levelOrder_list_pkg(root *tree.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	level := 0
	q := list.New()
	q.PushBack(root)

	for q.Len() > 0 {
		count := q.Len()
		res = append(res, []int{})
		for count > 0 {
			node := q.Remove(q.Front()).(*tree.TreeNode)
			res[level] = append(res[level], node.Val)
			count--
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
		level++
	}
	return res
}

func levelOrder(root *tree.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	level := 0
	q := queue.Queue{}
	q.Enqueue(root)

	for !q.IsEmpty() {
		count := q.Size()
		res = append(res, []int{})
		for count > 0 {
			dequeued, _ := q.Dequeue().(*tree.TreeNode)
			res[level] = append(res[level], dequeued.Val)
			count--
			if dequeued.Left != nil {
				q.Enqueue(dequeued.Left)
			}
			if dequeued.Right != nil {
				q.Enqueue(dequeued.Right)
			}
		}
		level++
	}
	return res
}
