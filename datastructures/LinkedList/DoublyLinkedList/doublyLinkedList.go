package doublyllist

import (
	"errors"

	linkedlist "github.com/go_problem_solving/datastructures/linkedList"
)

type Node struct {
	Value    interface{}
	Previous *Node
	Next     *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewDoublyLinkedList() linkedlist.LinkedList {
	return &DoublyLinkedList{}
}

func (ll *DoublyLinkedList) Insert(data interface{}) {
	newNode := ll.createNode(data)
	if ll.Head == nil {
		ll.Head = newNode
	} else {
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}
	ll.Size++
}

func (ll *DoublyLinkedList) Get(index int) (*Node, error) {
	if ll.Head == nil || index < 0 || index >= ll.Size {
		return nil, errors.New("Not found")
	}
	i := 0
	curr := ll.Head
	for {
		if index == i {
			return curr, nil
		}
		i++
		curr = curr.Next
	}

}

func (ll *DoublyLinkedList) Delete(data interface{}) (bool, error) {
	var prev, curr, nxt *Node
	prev = nil
	curr = ll.Head
	for {
		if curr.Value == data {
			prev.Next = curr.Next
			curr.Next.Previous = prev
		}
	}
}

func (ll *DoublyLinkedList) createNode(data interface{}) *Node {
	return &Node{
		Value: data,
	}
}
