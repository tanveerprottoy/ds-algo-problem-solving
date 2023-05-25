package stackqueue

import (
	"fmt"

	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/queue"
)

/*
LIFO ds
implemented using stacks

Method 1 (By making push operation costly)
This method makes sure that newest/last entered element is always at
the front of queue1, so that pop operation just dequeues from queue1.
To put the element at top of queue1, queue2 is used.

Method 2 (By making pop operation costly):
In this method, The new element is always enqueued to queue1.
In pop() operation, if queue2 is empty then all the elements 
except the last, are moved to q2. Finally, the last element is 
dequeued from q1 and returned
*/
type StackQueue[T any] struct {
	queue0 *queue.Queue[T]
	queue1 *queue.Queue[T]
	size   int
}

func NewStackQueue[T any]() *StackQueue[T] {
	return &StackQueue[T]{nil, nil, 0}
}

func (q *StackQueue[T]) Size() int {
	return q.size
}

func (s *StackQueue[T]) IsEmpty() bool {
	return s.size < 1
}

func (s *StackQueue[T]) PushHeavy(v T) {
	/*
		enqueue v to queue1.
		dequeue all items from queue0 and enqueue to queue1.
		Swap the queues of queue0 and queue1.
	*/
	s.queue1.Enqueue(v)
	for !s.queue0.IsEmpty() {
		v, err := s.queue0.Dequeue()
		if err != nil {
			fmt.Println(err)
			return
		}
		s.queue1.Enqueue(v)
	}
	// swap queueus
	q := s.queue0
	s.queue0 = s.queue1
	s.queue1 = q
}

func (s *StackQueue[T]) PopLight() (T, error) {
	// dequeue from queue0
	return s.queue0.Dequeue()
}

func (s *StackQueue[T]) PushLight(v T) {
	// enqueue to queue0
	s.queue0.Enqueue(v)
}

func (s *StackQueue[T]) PopHeavy() (T, error) {
	/*
		dequeue all items except the last from queue0 and
		enqueue to queue1.
		dequeue the last item from queue0, it's the result
		Swap the queues of queue0 and queue1.
	*/
	for s.queue0.Size() != 1 {
		v, err := s.queue0.Dequeue()
		if err != nil {
			return v, err
		}
		s.queue1.Enqueue(v)
	}
	var v T
	v, err := s.queue0.Dequeue()
	if err != nil {
		fmt.Println(err)
		return v, err
	}
	// swap queues
	q := s.queue0
	s.queue0 = s.queue1
	s.queue1 = q
	return v, nil
}
