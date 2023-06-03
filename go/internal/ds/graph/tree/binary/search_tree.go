package binary

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

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

/*
Allocate the memory for tree.
Set the data part to the value and set the left and right pointer of tree, point to NULL.
If the item to be inserted, will be the first element of the tree, then the left and right of this node will point to NULL.
Else, check if the item is less than the root element of the tree, if this is true, then recursively perform this operation with the left of the root.
If this is false, then perform this operation recursively with the right sub-tree of the root.
*/
func (t *SearchTree[T]) Insert(n *Node[T], v T) *Node[T] {
	/*
	 * It will handle two cases,
	 * 1. if the tree is empty, return new node in root
	 * 2. if the tree traversal reaches NULL, it will return the new node
	 */
	if n == nil {
		n = NewNode[T](v, nil, nil)
		return n
	}
	if n.val > v {
		/*
		* if given v is smaller than n.val,
		* need to find the correct place in left subtree and insert the new node
		 */
		n.left = t.Insert(n.left, v)
	} else if n.val < v {
		/*
		* if given v is greater than n.val,
		* need to find the correct place in right subtree and insert the new node
		 */
		n.right = t.Insert(n.right, v)
	}
	/*
	 * It will handle two cases
	 * (Prevent the duplicate nodes in the tree)
	 * 1.if n.val == v it will straight away return the address of the node
	 * 2.After the insertion, it will return the original unchanged node's address
	 */
	return n
}

func (t *SearchTree[T]) InsertIterative(v T) *Node[T] {
	n := NewNode[T](v, nil, nil)
	if t.Root == nil {
		t.Root = n
		return n
	}
	var prev *Node[T]
	curr := t.Root
	for curr != nil {
		if curr.val > v {
			// go left
			prev = curr
			curr = curr.left
		} else if curr.val < v {
			// go right
			prev = curr
			curr = curr.right
		}
	}
	if prev.val > v {
		prev.left = n
	} else {
		prev.right = n
	}
	return n
}

func (t *SearchTree[T]) Search(n *Node[T], v T) *Node[T] {
	if n == nil || n.val == v {
		return n
	}
	if n.val < v {
		return t.Search(n.right, v)
	}
	// n.val > v
	return t.Search(n.left, v)
}

func (t *SearchTree[T]) SearchIterative(v T) *Node[T] {
	if t.Root == nil {
		return nil
	}
	curr := t.Root
	for curr != nil {
		if curr.val < v {
			curr = curr.left
		} else if curr.val > v {
			curr = curr.right
		} else {
			// curr.val == v
			return curr
		}
	}
	return nil
}

/*
left->root->right
*/
func (t *SearchTree[T]) InorderTraverse(n *Node[T]) {
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
func (t *SearchTree[T]) PreorderTraverse(n *Node[T]) {
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
func (t *SearchTree[T]) PostorderTraverse(n *Node[T]) {
	if n == nil {
		return
	}
	t.PostorderTraverse(n.left)
	t.PostorderTraverse(n.right)
	fmt.Println(n.val)
}
