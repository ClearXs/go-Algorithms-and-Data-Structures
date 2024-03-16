package queue

// Queue 双端队列
type Queue[T interface{}] interface {
	Push(ele *T)
	// Offer 不删除元素，返回队列尾元素
	Offer() *T
	Pop() *T
	Peek(index int) *T
	PeekLast() *T
	Size() int
	Empty() bool
}

type Impl[T interface{}] struct {
	items []*T
}

func (queue *Impl[T]) Empty() bool {
	return queue.Size() == 0
}

// Push enqueue
func (queue *Impl[T]) Push(ele *T) {
	queue.items = append(queue.items, ele)
}

func (queue *Impl[T]) Offer() *T {
	return queue.Peek(0)
}

// Pop dequeue return generic ele or nil if queue is empty
func (queue *Impl[T]) Pop() *T {
	if queue.Size() == 0 {
		return nil
	}
	ele := queue.items[0]
	queue.items = queue.items[1:]
	return ele
}

// Peek peek queue element for index
func (queue *Impl[T]) Peek(index int) *T {
	return queue.items[index]
}

// PeekLast peek stack element for last
func (queue *Impl[T]) PeekLast() *T {
	size := queue.Size()
	if size == 0 {
		return nil
	}
	return queue.items[size-1]
}

func (queue *Impl[T]) Size() int {
	return len(queue.items)
}

func NewQueue[T interface{}]() Queue[T] {
	return &Impl[T]{items: make([]*T, 0)}
}
