package queue

import "errors"

/*
	FIFO ds
*/
type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{make([]T, 0)}
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() < 1
}

func (q *Queue[T]) shift() {
	// shift the elements
	// to 1 index earlier
	for i := 1; i < len(q.data); i++ {
		q.data[i-1] = q.data[i]
	}
}

func (q *Queue[T]) Enqueue(e T) {
	// insert to last
	q.data = append(q.data, e)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var e T
	if q.IsEmpty() {
		return e, errors.New("queue empty")
	}
	e = q.data[0]
	q.shift()
	return e, nil
}
