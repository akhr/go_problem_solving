package singlyllist

import "errors"

// Node : Linked list node
type Node struct {
	Next  *Node
	Value interface{}
}

// SinglyLinkedList :
type SinglyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

// NewSinglyLinkedList :
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (ll *SinglyLinkedList) createNode(data interface{}) *Node {
	return &Node{
		Value: data,
	}
}

// Insert : Insert into/from SinglyLinkedList
func (ll *SinglyLinkedList) Insert(data interface{}) {
	newNode := ll.createNode(data)
	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}
	ll.Size++
}

// Get : Get into/from SinglyLinkedList
func (ll *SinglyLinkedList) Get(index int) (interface{}, error) {
	if index >= ll.Size {
		return nil, errors.New("Not Found")
	}
	curr := ll.Head
	i := 0
	for {
		if i == index {
			return curr.Value, nil
		}
		curr = curr.Next
		i++
	}
}

// Delete : Delete into/from SinglyLinkedList
func (ll *SinglyLinkedList) Delete(data interface{}) (bool, error) {
	if ll.Head.Value == data {
		ll.Head = ll.Head.Next
		ll.Size--
		return true, nil
	}
	prev := ll.Head
	curr := ll.Head.Next
	for {
		if curr != nil && curr.Value == data {
			prev.Next = curr.Next
			ll.Size--
			return true, nil
		}
		if curr.Next == nil {
			return false, errors.New("Not Found")
		}
		prev = curr
		curr = curr.Next
	}
}
