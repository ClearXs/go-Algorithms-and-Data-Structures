package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	queue := NewMinimumPriorityQueue[int](func(t1 *int, t2 *int) int {
		return *t1 - *t2
	})
	i1 := 2
	i2 := 1
	queue.Push(&i1)
	queue.Push(&i2)

	assert.Equal(t, 2, queue.Size())

	o1 := queue.Offer()
	assert.Equal(t, 1, *o1)

	last := queue.PeekLast()
	assert.Equal(t, 2, *last)

	p1 := queue.Pop()
	p2 := queue.Pop()
	assert.Equal(t, 1, *p1)
	assert.Equal(t, 2, *p2)
}

func TestPriorityQueueIfEmpty(t *testing.T) {
	queue := NewMinimumPriorityQueue[int](func(t1 *int, t2 *int) int {
		return *t1 - *t2
	})

	pop := queue.Pop()
	assert.Empty(t, pop)

	size := queue.Size()
	assert.Equal(t, 0, size)

	last := queue.PeekLast()
	assert.Empty(t, last)

	offer := queue.Offer()
	assert.Empty(t, offer)
}
