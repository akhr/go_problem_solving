package symmetricTree101

import (
	"container/list"

	tree "github.com/go_problem_solving/leetcode/Tree"
)

func isSymmetric(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root, root)
}

func isMirror(t1, t2 *tree.TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return isMirror(t1.Left, t2.Right) &&
		isMirror(t1.Right, t2.Left)
}

func isSymmetric_2(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	q := list.New()
	q.PushBack(root)
	q.PushBack(root)

	for q.Len() > 0 {
		t1 := q.Remove(q.Front()).(*tree.TreeNode)
		t2 := q.Remove(q.Front()).(*tree.TreeNode)
		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t1.Val != t2.Val {
			return false
		}
		q.PushBack(t1.Left)
		q.PushBack(t2.Right)
		q.PushBack(t1.Right)
		q.PushBack(t2.Left)
	}
	return true
}
