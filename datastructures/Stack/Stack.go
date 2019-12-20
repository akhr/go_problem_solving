package stack

type Stack struct {
	data []interface{}
	size int
}

func (s *Stack) Push(item interface{}) {
	if item != nil {
		s.data = append(s.data, item)
		s.size++
	}
}

func (s *Stack) Pop() interface{} {
	var item interface{}
	if s.size > 0 {
		item = s.data[s.size-1]
		s.data = s.data[:s.size-1]
		s.size--
	}
	return item
}

func (s *Stack) Peek() interface{} {
	return s.data[s.size-1]
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) GetSize() int {
	return s.size
}
