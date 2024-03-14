package queue

import "data-structure/heap"

type PriorityQueue[T interface{}] struct {
	heap *heap.Heap[T]
}

func (p *PriorityQueue[T]) Push(ele *T) {
	p.heap.Insert(ele)
}

func (p PriorityQueue[T]) Offer() *T {
	return p.heap.Peek(1)
}

func (p PriorityQueue[T]) Pop() *T {
	return p.heap.DeleteExtreme()
}

func (p PriorityQueue[T]) Peek(index int) *T {
	return p.heap.Peek(index)
}

func (p PriorityQueue[T]) PeekLast() *T {
	return p.heap.Peek(p.Size())
}

func (p PriorityQueue[T]) Size() int {
	return p.heap.Size()
}

func NewPriorityQueue[T interface{}](sort heap.SortSymbolism, comparator func(t1 *T, t2 *T) int) Queue[T] {
	return &PriorityQueue[T]{heap: heap.NewHeap[T](sort, comparator)}
}

func NewMaximumPriorityQueue[T interface{}](comparator func(t1 *T, t2 *T) int) Queue[T] {
	return &PriorityQueue[T]{heap: heap.NewMaximumHeap[T](comparator)}
}

func NewMinimumPriorityQueue[T interface{}](comparator func(t1 *T, t2 *T) int) Queue[T] {
	return &PriorityQueue[T]{heap: heap.NewMinimumHeap[T](comparator)}
}
