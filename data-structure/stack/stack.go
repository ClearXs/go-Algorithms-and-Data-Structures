package stack

type Stack[T interface{}] struct {
	items []T
}

// Push 入栈
func (s *Stack[T]) Push(ele T) {
	s.items = append(s.items, ele)
}

// Pop 出栈
func (s *Stack[T]) Pop() *T {
	if len(s.items) == 0 {
		return nil
	}
	size := len(s.items)
	ele := s.items[size-1]
	s.items = s.items[:size-1]
	return &ele
}

// Peek peek stack element for index
func (s *Stack[T]) Peek(index int) *T {
	return &s.items[index]
}

// PeekLast peek stack element for last
func (s *Stack[T]) PeekLast() *T {
	size := s.Size()
	if size == 0 {
		return nil
	}
	return &s.items[size-1]
}

// GetItems return intrinsic items
func (s *Stack[T]) GetItems() []T {
	return s.items
}

// Size return stack size
func (s *Stack[T]) Size() int {
	return len(s.items)
}

func New[T interface{}]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}
