package singlyllist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList_Create(t *testing.T) {
	SinglyLinkedList := NewSinglyLinkedList()
	assert.NotNil(t, SinglyLinkedList)
	assert.Nil(t, SinglyLinkedList.Head)
	assert.Nil(t, SinglyLinkedList.Tail)
}

func TestSinglyLinkedList_Insert(t *testing.T) {
	SinglyLinkedList := NewSinglyLinkedList()

	SinglyLinkedList.Insert(100)
	SinglyLinkedList.Insert(200)
	SinglyLinkedList.Insert(300)
	SinglyLinkedList.Insert(400)

	item, err := SinglyLinkedList.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 100, item)

	item, err = SinglyLinkedList.Get(3)
	assert.NoError(t, err)
	assert.Equal(t, 400, item)

	item, err = SinglyLinkedList.Get(200)
	assert.Error(t, err)
	assert.Equal(t, "Not Found", err.Error())
}

func TestSinglyLinkedList_Delete(t *testing.T) {
	SinglyLinkedList := NewSinglyLinkedList()

	SinglyLinkedList.Insert(100)
	SinglyLinkedList.Insert(200)
	SinglyLinkedList.Insert(300)
	SinglyLinkedList.Insert(400)
	SinglyLinkedList.Delete(400) //tail
	SinglyLinkedList.Delete(200) //something in the middle
	SinglyLinkedList.Delete(100) //head

	item, err := SinglyLinkedList.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 300, item)

	item, err = SinglyLinkedList.Get(2)
	assert.Error(t, err)
	assert.Equal(t, "Not Found", err.Error())
}
