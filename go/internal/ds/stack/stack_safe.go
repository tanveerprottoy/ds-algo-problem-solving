package stack

import (
	"errors"
	"sync"
)

// thread safe with lock
type StackSafe[T any] struct {
	lock sync.Mutex
	data []T
}

func NewStackSafe[T any]() *StackSafe[T] {
	return &StackSafe[T]{sync.Mutex{}, make([]T, 0)}
}

func (s *StackSafe[T]) IsEmpty() bool {
	return len(s.data) < 1
}

func (s *StackSafe[T]) Length() int {
	return len(s.data)
}

func (s *StackSafe[T]) Push(v T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = append(s.data, v)
}

func (s *StackSafe[T]) Pop() (T, error) {
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

func (s *StackSafe[T]) PopAlt() T {
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
