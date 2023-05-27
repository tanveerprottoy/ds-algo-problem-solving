package binary

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

/*
Insert − Inserts data in a tree.
Preorder Traversal – perform Traveling a tree in a pre-order manner in data structure .
In order Traversal – perform Traveling a tree in an in-order manner.
Post order Traversal –perform Traveling a tree in a post-order manner.
*/
type Tree[T constraints.Ordered] struct {
	Root *Node[T]
}

func NewTree[T constraints.Ordered](root *Node[T]) *Tree[T] {
	return &Tree[T]{Root: root}
}

func (t *Tree[T]) Insert(v T) {
	curr := t.Root
	for curr != nil {
		if curr.val < v {
			// go to left
			curr = curr.left
		} else {
			// go to right
			curr = curr.right
		}
	}
	// n := NewNode(v, nil, nil)
}

/*
left->root->right
*/
func (t *Tree[T]) InorderTraverse(n *Node[T]) {
	if n == nil {
		return
	}
	t.InorderTraverse(n.left)
	fmt.Println(n.val)
	t.InorderTraverse(n.right)
}

/*
root->left->right
*/
func (t *Tree[T]) PreorderTraverse(n *Node[T]) {
	if n == nil {
		return
	}
	fmt.Println(n.val)
	t.PreorderTraverse(n.left)
	t.PreorderTraverse(n.right)
}

/*
left->right->root
*/
func (t *Tree[T]) PostorderTraverse(n *Node[T]) {
	if n == nil {
		return
	}
	t.PreorderTraverse(n.left)
	t.PreorderTraverse(n.right)
	fmt.Println(n.val)
}
