package stack

import (
	"errors"
)

// not thread safe
type StackQueue[T any] struct {
	data []T
}

func NewStackQueue[T any]() *StackQueue[T] {
	return &StackQueue[T]{make([]T, 0)}
}

func (s *StackQueue[T]) IsEmpty() bool {
	return len(s.data) < 1
}

func (s *StackQueue[T]) Length() int {
	return len(s.data)
}

func (s *StackQueue[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *StackQueue[T]) Pop() (T, error) {
	var res T
	if s.IsEmpty() {
		return res, errors.New("empty StackUnsafe")
	}
	l := s.Length()
	res = s.data[l-1]
	s.data = s.data[:l-1]
	return res, nil
}

func (s *StackQueue[T]) PopAlt() T {
	var res T
	if s.IsEmpty() {
		return res
	}
	l := s.Length()
	res = s.data[l-1]
	s.data = s.data[:l-1]
	return res
}

func (s *StackQueue[T]) Peek() T {
	var d T
	if s.IsEmpty() {
		return d
	}
	// return top/last element
	d = s.data[len(s.data)-1]
	return d
}
