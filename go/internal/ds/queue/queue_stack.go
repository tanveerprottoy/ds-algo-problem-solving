package queue

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
	// pop from stack0 and store popped item in stack1
	// push the new item in stack0
	// push back items from stack1 to stack0
	for !q.stack0.IsEmpty() {
		p, err := q.stack0.Pop()
		if err != nil {
			fmt.Errorf(err.Error())
		}
		q.stack1.Push(p)
	}
	q.stack0.Push(v)
	q.size++
	for !q.stack1.IsEmpty() {
		p, err := q.stack1.Pop()
		if err != nil {
			fmt.Errorf(err.Error())
		}
		q.stack0.Push(p)
	}
}

func (q *QueueStack[T]) DequeueLight() (T, error) {
	// with enqueueHeavy DequeueLight is simple
	// as the oldest/first item is already on top
	// of stack popping it will do it
	p, err := q.stack0.Pop()
	if err != nil {
		return p, err
	}
	return p, nil
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
	var p T
	if q.stack0.IsEmpty() && q.stack1.IsEmpty() {
		return p, errors.New("both stacks are empty")
	}
	if q.stack1.IsEmpty() {
		for !q.stack0.IsEmpty() {
			p, err := q.stack0.Pop()
			if err != nil {
				return p, err
			}
			q.stack1.Push(p)
		}
	}
	p, err := q.stack1.Pop()
	q.size++
	return p, err
}
