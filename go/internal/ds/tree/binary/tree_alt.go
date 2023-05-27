package binary

import (
	"fmt"

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

func (t *TreeAlt[T]) InorderTraverseStack() {
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
