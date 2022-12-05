package main


type item[T any] struct {
	v T
	next *item[T]
}

func (i *item[T]) value() T {
	return i.v
}

// not save for concurrent use
type stack[T any] struct {
	head *item[T]
}

func (s *stack[T]) peek() T {
	if s.head == nil {
		panic("peek on empty stack")
	}
	return s.head.value()
}

func (s *stack[T]) pop() T {
	if s.head == nil {
		panic("pop on empty stack")
	}
	head := s.head
	s.head = head.next
	return head.value()
}

func (s *stack[T]) popn(n int) []T {
	x := make([]T, n, n)
	for i := range x {
		x[i] = s.pop()
	}
	return x
}

func (s *stack[T]) push(v T) {
	s.head = &item[T]{v, s.head}
}

func (s *stack[T]) isEmpty() bool {
	if s.head == nil {
		return true
	}
	return false
}
