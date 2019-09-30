package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fib(t *testing.T) {
	in := 4
	expectedSlice := []int{0, 1, 1, 2}
	actualSlice := fib(in)
	t.Log("actualSize ", actualSlice)
	assert.ElementsMatch(t, expectedSlice, actualSlice)
}
