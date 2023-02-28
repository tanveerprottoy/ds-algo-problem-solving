package ds

import (
	"errors"
	"sync"
)

// thread safe with lock
type Stack[T any] struct {
	lock sync.Mutex
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{sync.Mutex{}, make([]T, 0)}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Length() int {
	return len(s.data)
}

func (s *Stack[T]) Push(v T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, error) {
	var res T
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.IsEmpty() {
		return res, errors.New("empty stack")
	}
	l := s.Length()
	res = s.data[l-1]
	s.data = s.data[:l-1]
	return res, nil
}

func (s *Stack[T]) PopAlt() T {
	var res T
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.IsEmpty() {
		return res
	}
	l := s.Length()
	res = s.data[l-1]
	s.data = s.data[:l-1]
	return res
}
