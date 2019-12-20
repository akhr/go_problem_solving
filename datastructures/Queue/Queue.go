package queue

type Queue struct {
	data []interface{}
	size int
}

func (q *Queue) Enqueue(item interface{}) {
	q.data = append(q.data, item)
	q.size++
}

func (q *Queue) Dequeue() interface{} {
	var item interface{}
	if q.size > 0 {
		item = q.data[0]
		q.data = q.data[1:]
		q.size--
	}
	return item
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
