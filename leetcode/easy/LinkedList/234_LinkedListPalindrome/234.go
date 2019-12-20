package linkedlistpalindrome234

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	endOfFirstHalf := slow
	l1 := head
	l2 := reverse(endOfFirstHalf.Next)
	result := compare(l1, l2)

	//Restore original list
	endOfFirstHalf.Next = reverse(l2)
	return result
}

func print(node *ListNode) {
	fmt.Println("Start Printing...")
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func compare(l1, l2 *ListNode) bool {
	for ; l2 != nil; l1, l2 = l1.Next, l2.Next {
		if l1.Val != l2.Val {
			return false
		}
	}
	return true
}
