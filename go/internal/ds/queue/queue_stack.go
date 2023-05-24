package queue

import (
	"errors"

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
	stack0 *stack.Stack[int]
	stack1 *stack.Stack[int]
	size   int
}

func NewQueueStack[T any]() *QueueStack[T] {
	return &QueueStack[T]{nil, nil, 0}
}

func (q *QueueStack[T]) Size() int {
	return q.size
}

func (q *QueueStack[T]) IsEmpty() bool {
	return q.Size() < 1
}

func (q *QueueStack[T]) EnqueueHeavy(v T) {
	if 
}

func (q *QueueStack[T]) DequeueLight() (T, error) {
	
}

func (q *QueueStack[T]) EnqueueLight(v T) {
	// insert to last
	
}

func (q *QueueStack[T]) DequeueHeavy() (T, error) {
	
}
