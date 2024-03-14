package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := New[string]()
	stack.Push("1")
	s := stack.Pop()
	assert.Equal(t, "1", *s)

	stack.Push("2")
	stack.Push("3")
	assert.Equal(t, "3", *stack.Pop())
	assert.Equal(t, "2", *stack.Pop())

	assert.Equal(t, 0, stack.Size())

	assert.Empty(t, stack.Pop())
}
