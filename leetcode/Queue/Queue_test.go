package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	q := Queue{}
	assert.True(t, q.IsEmpty())
	q.Enqueue(1)
	assert.False(t, q.IsEmpty())
	q.Dequeue()
	assert.True(t, q.IsEmpty())
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	t.Logf("\nQ in Step A - %#v", q)
	assert.Equal(t, 3, q.Size())
	q.Dequeue()
	q.Dequeue()
	assert.Equal(t, 4, q.Dequeue())
	q.Dequeue() // Already stack is empty
	t.Logf("\nQ in Step B - %#v", q)
	q.Enqueue(20)
	t.Logf("\nQ in Step C - %#v", q)
	q.Enqueue("30")
	t.Logf("\nQ in Step D - %#v", q)
	assert.Equal(t, 20, q.Dequeue())
	assert.Equal(t, "30", q.Dequeue())
}
