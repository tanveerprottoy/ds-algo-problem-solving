package binary

import (
	"fmt"

	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/stack"
	"golang.org/x/exp/constraints"
)

/*
Insert − Inserts data in a TreeAlt.
Preorder Traversal – perform Traveling a TreeAlt in a pre-order manner in data structure .
In order Traversal – perform Traveling a TreeAlt in an in-order manner.
Post order Traversal –perform Traveling a TreeAlt in a post-order manner.
*/
type TreeAlt[T constraints.Ordered] struct {
	Root  *Node[T]
	count int
}

func NewTreeAlt[T constraints.Ordered](root *Node[T]) *TreeAlt[T] {
	return &TreeAlt[T]{Root: root, count: 0}
}

func (t *TreeAlt[T]) Insert(v T) {
	var curr *Node[T]
	if t.count%2 == 0 {
		// for even go left
		curr = t.Root.left
		for curr.left != nil {
			curr = curr.left
		}
		n := NewNode(v, nil, nil)
		curr.left = n
	} else {
		// for odd go right
		curr = t.Root.right
		for curr.right != nil {
			curr = curr.right
		}
		n := NewNode(v, nil, nil)
		curr.right = n
	}
}

/*
left->root->right
*/
func (t *TreeAlt[T]) InorderTraverse() {
	
}

/*

Create an empty stack
Initialize the current node as root.
Push the current node to S and set current = current->left 
until current is NULL
If current is NULL and the stack is not empty then:
Pop the top item from the stack.
Print the popped item and set current = popped_item->right 
Go to step 3.
If current is NULL and the stack is empty then we are done.

Start from the root, call it PTR.
Push PTR onto stack if PTR is not NULL.
Move to left of PTR and repeat step 2.
If PTR is NULL and stack is not empty, 
then Pop element from stack and set as PTR.
Process PTR and move to right of PTR , go to step 2.
*/
func (t *TreeAlt[T]) InorderTraverseStack() {
	s := stack.
	r := t.Root
	if r != nil {

	}
	fmt.Println("InorderTraverseStack")
}

/*
root->left->right
*/
func (t *TreeAlt[T]) PreorderTraverse() {
}

/*
left->right->root
*/
func (t *TreeAlt[T]) PostorderTraverse() *Node[T] {
	return nil
}
