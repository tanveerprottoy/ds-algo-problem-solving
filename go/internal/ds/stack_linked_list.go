package ds

import (
	"errors"
	"sync"
)

type StackLinkedList[T any] struct {
	lock sync.Mutex
	data []T
}

func NewStackLinkedList[T any]() *Stack[T] {
	return &Stack[T]{sync.Mutex{}, make([]T, 0)}
}

func (s *StackLinkedList[T]) Push(v T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = append(s.data, v)
}

func (s *StackLinkedList[T]) Pop() (T, error) {
	var res T
	s.lock.Lock()
	defer s.lock.Unlock()
	l := len(s.data)
	if l == 0 {
		return res, errors.New("empty stack")
	}
	res = s.data[l-1]
	s.data = s.data[:l-1]
	return res, nil
}
