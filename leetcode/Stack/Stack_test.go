package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := Stack{}
	assert.True(t, s.IsEmpty())
	s.Push(1)
	assert.False(t, s.IsEmpty())
	s.Pop()
	assert.True(t, s.IsEmpty())
	s.Push(2)
	s.Push(3)
	s.Push(4)
	t.Logf("\nStack in Step A - %#v", s)
	assert.Equal(t, 3, s.GetSize())
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop() // Already stack is empty
	t.Logf("\nStack in Step B - %#v", s)
	s.Push(20)
	t.Logf("\nStack in Step C - %#v", s)
	s.Push("30")
	t.Logf("\nStack in Step D - %#v", s)
}
