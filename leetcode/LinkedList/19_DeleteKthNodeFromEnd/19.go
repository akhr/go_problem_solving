package deletekthnodefromend19

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	preHead := &ListNode{Next: head}
	slow := preHead
	fast := head
	for pos := n; pos > 0; pos-- {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return preHead.Next
}
