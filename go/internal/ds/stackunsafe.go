package ds

import (
	"errors"
)

// not thread safe
type StackUnsafe[T any] struct {
	data []T
}

func NewUnsafeStack[T any]() *StackUnsafe[T] {
	return &StackUnsafe[T]{make([]T, 0)}
}

func (s *StackUnsafe[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *StackUnsafe[T]) Length() int {
	return len(s.data)
}

func (s *StackUnsafe[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *StackUnsafe[T]) Pop() (T, error) {
	var res T
	if s.IsEmpty() {
		return res, errors.New("empty StackUnsafe")
	}
	l := s.Length()
	res = s.data[l-1]
	s.data = s.data[:l-1]
	return res, nil
}

func (s *StackUnsafe[T]) PopAlt() T {
	var res T
	if s.IsEmpty() {
		return res
	}
	l := s.Length()
	res = s.data[l-1]
	s.data = s.data[:l-1]
	return res
}
