package stack

import (
	"errors"
)

type StackUnsafeNew[T any] struct {
	data []T
}

func NewStackUnsafeNew[T any]() *StackUnsafeNew[T] {
	return &StackUnsafeNew[T]{make([]T, 0)}
}

func (s *StackUnsafeNew[T]) Push(e T) {
	s.data = append(s.data, e)
}

func (s *StackUnsafeNew[T]) Pop() (T, error) {
	var e T
	l := len(s.data)
	if l < 1 {
		return e, errors.New("stack is empty")
	}
	e = s.data[l-1]
	s.data = s.data[:l-1]
	return e, nil
}

func (s *StackUnsafeNew[T]) Peek() T {
	// return top/last element
	return s.data[len(s.data)-1]
}
