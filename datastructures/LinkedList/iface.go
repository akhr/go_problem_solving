package linkedlist

type LinkedList interface {
	Insert(data interface{})
	Get(index int) (*Node, error)
	Delete(data interface{}) (bool, error)
}
