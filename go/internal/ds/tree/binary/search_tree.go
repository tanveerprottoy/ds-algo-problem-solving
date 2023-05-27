package binary

import "golang.org/x/exp/constraints"

/*
Insert − Inserts data in a SearchTree.
Search − Searches specific data  in a SearchTree to check it is present or not.
Preorder Traversal – perform Traveling a SearchTree in a pre-order manner in data structure .
In order Traversal – perform Traveling a SearchTree in an in-order manner.
Post order Traversal –perform Traveling a SearchTree in a post-order manner.
*/
type SearchTree[T constraints.Ordered] struct {
	Root *Node[T]
}

func NewSearchTree[T constraints.Ordered](root *Node[T]) *SearchTree[T] {
	return &SearchTree[T]{Root: root}
}

func (t *SearchTree[T]) Insert(v T) {
	curr := t.Root.left
	for curr != nil {
		if curr.val <= v {
			// go left
			curr = curr.left
		} else {
			// go right
			curr = curr.right
		}
	}
}

func (t *SearchTree[T]) Search(v T) *Node[T] {
	return nil
}

/*
left->root->right
*/
func (t *SearchTree[T]) InorderTraverse() *Node[T] {
	return nil
}

/*
root->left->right
*/
func (t *SearchTree[T]) PreorderTraverse() *Node[T] {
	return nil
}

/*
left->right->root
*/
func (t *SearchTree[T]) PostorderTraverse() *Node[T] {
	return nil
}
