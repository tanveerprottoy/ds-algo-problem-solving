package stack

import (
	"errors"
)

// not thread safe
type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{make([]T, 0)}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) < 1
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, error) {
	var res T
	if s.IsEmpty() {
		return res, errors.New("empty StackUnsafe")
	}
	l := s.Size()
	res = s.data[l-1]
	s.data = s.data[:l-1]
	return res, nil
}

func (s *Stack[T]) PopAlt() T {
	var res T
	if s.IsEmpty() {
		return res
	}
	l := s.Size()
	res = s.data[l-1]
	s.data = s.data[:l-1]
	return res
}

func (s *Stack[T]) Peek() T {
	var d T
	if s.IsEmpty() {
		return d
	}
	// return top/last element
	d = s.data[len(s.data)-1]
	return d
}
