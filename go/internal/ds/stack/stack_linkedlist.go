package stack

import (
	"errors"
	"sync"
)

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func NewNode[T any](val T, Next *Node[T]) *Node[T] {
	return &Node[T]{Val: val, Next: Next}
}

type StackLinkedList[T any] struct {
	lock   sync.Mutex
	top    *Node[T]
	length int
}

func NewStackLinkedList[T any]() *StackSafe[T] {
	return &StackSafe[T]{sync.Mutex{}, nil}
}

func (s *StackLinkedList[T]) IsEmpty() bool {
	return s.top == nil
}

func (s *StackLinkedList[T]) Length() int {
	return s.length
}

func (s *StackLinkedList[T]) Push(v T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.IsEmpty() {
		n := NewNode[T](v, nil)
		s.top = n
		s.length++
		return
	}
	n := NewNode[T](v, s.top)
	s.top = n
	s.length++
}

func (s *StackLinkedList[T]) Pop() (T, error) {
	var res T
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.IsEmpty() {
		return res, errors.New("empty stack")
	}
	res = s.top.Val
	s.top = s.top.Next
	s.length--
	return res, nil
}

func (s *StackLinkedList[T]) Peek() T {
	var d T
	if s.IsEmpty() {
		return d
	}
	// return top/last element
	d = s.top.Val
	return d
}
