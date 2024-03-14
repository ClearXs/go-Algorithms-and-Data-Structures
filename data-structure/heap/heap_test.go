package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	heap := NewMinimumHeap[int](func(t1 *int, t2 *int) int {
		return *t1 - *t2
	})
	i1 := 2
	i2 := 1
	heap.Insert(&i1)
	heap.Insert(&i2)
	assert.Equal(t, 1, *heap.Peek(1))
	assert.Equal(t, 2, *heap.Peek(2))
	i3 := 5
	i4 := 4
	heap.Insert(&i3)
	heap.Insert(&i4)
	assert.Equal(t, 5, *heap.Peek(3))
	assert.Equal(t, 4, *heap.Peek(4))
	i5 := 3
	heap.Insert(&i5)
	assert.Equal(t, 3, *heap.Peek(5))

	i6 := 1
	heap.Insert(&i6)
	assert.Equal(t, 1, *heap.Peek(3))
	assert.Equal(t, 5, *heap.Peek(6))
}

func TestDelete(t *testing.T) {
	heap := NewMinimumHeap[int](func(t1 *int, t2 *int) int {
		return *t1 - *t2
	})
	i1 := 2
	i2 := 1
	i3 := 5
	i4 := 4
	i5 := 3
	i6 := 1
	heap.InsertArray([]*int{&i1, &i2, &i3, &i4, &i5, &i6})

	e1 := heap.DeleteExtreme()
	assert.Equal(t, 1, *e1)

	e2 := heap.DeleteExtreme()
	assert.Equal(t, 1, *e2)

	e3 := heap.DeleteExtreme()
	assert.Equal(t, 2, *e3)

	e4 := heap.DeleteExtreme()
	assert.Equal(t, 3, *e4)

	e5 := heap.DeleteExtreme()
	assert.Equal(t, 4, *e5)

	e6 := heap.DeleteExtreme()
	assert.Equal(t, 5, *e6)

	e7 := heap.DeleteExtreme()
	assert.Empty(t, e7)
}
