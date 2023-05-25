package queue

import "errors"

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func NewNode[T any](val T, Next *Node[T]) *Node[T] {
	return &Node[T]{Val: val, Next: Next}
}

/*
	FIFO ds
*/
type QueueLinkedList[T any] struct {
	front *Node[T]
	rear  *Node[T]
	size  int
}

func NewQueueLinkedList[T any](front, rear *Node[T]) *QueueLinkedList[T] {
	return &QueueLinkedList[T]{front, rear, 0}
}

func (q *QueueLinkedList[T]) Size() int {
	return q.size
}

func (q *QueueLinkedList[T]) IsEmpty() bool {
	return q.Size() < 1
}

func (q *QueueLinkedList[T]) Enqueue(v T) {
	// insert to last
	if q.rear == nil {
		n := NewNode(v, nil)
		q.rear = n
		q.size++
		if q.front == nil {
			q.front = n
		}
		return
	}
	tmp := q.rear
	n := NewNode(v, nil)
	tmp.Next = n
	q.rear = n
	q.size++
}

func (q *QueueLinkedList[T]) Dequeue() (T, error) {
	var v T
	if q.IsEmpty() {
		return v, errors.New("queue empty")
	}
	v = q.front.Val
	q.front = q.front.Next
	return v, nil
}

func (q *QueueLinkedList[T]) Peek() (T, error) {
	var v T
	if q.IsEmpty() {
		return v, errors.New("queue empty")
	}
	v = q.front.Val
	return v, nil
}
