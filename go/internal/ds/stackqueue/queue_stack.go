package stackqueue

import (
	"errors"
	"fmt"

	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/stack"
)

/*
FIFO ds
implemented using stacks

Method 1 (By making enQueue operation costly)
This method makes sure that oldest entered element is always at
the top of stack 1, so that deQueue operation just pops from stack1.
To put the element at top of stack1, stack2 is used.

Method 2 (By making deQueue operation costly):
In this method, in en-queue operation, the new element is
entered at the top of stack1. In de-queue operation, if
stack2 is empty then all the elements are moved to stack2
and finally top of stack2 is returned.
*/
type QueueStack[T any] struct {
	stack0 *stack.Stack[T]
	stack1 *stack.Stack[T]
	size   int
}

func NewQueueStack[T any]() *QueueStack[T] {
	return &QueueStack[T]{nil, nil, 0}
}

func (q *QueueStack[T]) Size() int {
	return q.size
}

/*
to make it fifo, will have to store the oldest
item on the top with the enqueue heavy approach
*/
func (q *QueueStack[T]) EnqueueHeavy(v T) {
	// pop all items from stack0 and store popped items in stack1
	// push the new item in stack0
	// push back items from stack1 to stack0
	for !q.stack0.IsEmpty() {
		v, err := q.stack0.Pop()
		if err != nil {
			fmt.Println(err.Error())
		}
		q.stack1.Push(v)
	}
	q.stack0.Push(v)
	q.size++
	for !q.stack1.IsEmpty() {
		v, err := q.stack1.Pop()
		if err != nil {
			fmt.Println(err.Error())
		}
		q.stack0.Push(v)
	}
}

func (q *QueueStack[T]) DequeueLight() (T, error) {
	// with enqueueHeavy DequeueLight is simple
	// as the oldest/first item is already on top
	// of stack popping it will do it
	v, err := q.stack0.Pop()
	if err != nil {
		return v, err
	}
	return v, nil
}

func (q *QueueStack[T]) EnqueueLight(v T) {
	// push to stack0
	q.stack0.Push(v)
}

func (q *QueueStack[T]) DequeueHeavy() (T, error) {
	/*
			1) If both stacks are empty then error.
		  	2) If stack2 is empty
		       While stack1 is not empty, push everything from stack1 to stack2.
		  	3) Pop the element from stack2 and return it.
	*/
	var v T
	if q.stack0.IsEmpty() && q.stack1.IsEmpty() {
		return v, errors.New("both stacks are empty")
	}
	if q.stack1.IsEmpty() {
		for !q.stack0.IsEmpty() {
			v, err := q.stack0.Pop()
			if err != nil {
				return v, err
			}
			q.stack1.Push(v)
		}
	}
	v, err := q.stack1.Pop()
	q.size++
	return v, err
}

func (q *QueueStack[T]) Peek() (T, error) {
	var v T
	if q.Size() == 0 {
		return v, errors.New("queue empty")
	}
	v = q.stack0.Peek()
	return v, nil
}
