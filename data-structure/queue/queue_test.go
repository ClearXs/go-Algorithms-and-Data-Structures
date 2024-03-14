package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue[string]()

	s1 := "1"
	s2 := "2"
	queue.Push(&s1)
	queue.Push(&s2)

	assert.Equal(t, "1", *queue.Pop())
	assert.Equal(t, "2", *queue.Pop())

	s3 := "3"
	queue.Push(&s3)
	assert.Equal(t, "3", *queue.Pop())

	assert.Empty(t, queue.Pop())
}
