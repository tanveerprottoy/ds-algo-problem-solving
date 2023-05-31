package binary

import (
	"fmt"

	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/stack"
	"golang.org/x/exp/constraints"
)

/*
Insert − Inserts data in a tree.
Preorder Traversal – perform Traveling a tree in a pre-order manner in data structure .
In order Traversal – perform Traveling a tree in an in-order manner.
Post order Traversal –perform Traveling a tree in a post-order manner.
*/
type Tree[T constraints.Ordered] struct {
	Root  *Node[T]
	count int
}

func NewTree[T constraints.Ordered](root *Node[T]) *Tree[T] {
	return &Tree[T]{Root: root, count: 0}
}

func (t *Tree[T]) Insert(v T) {
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
	t.count++
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
1 Start from the root.
2 Push curr onto stack if curr is not NULL.
3 Move to left of curr and repeat step 2.
4 If curr is NULL and stack is not empty, then Pop element from stack and set as curr.
5 Process curr and move to right of curr , go to step 2
*/
func (t *Tree[T]) InorderTraverseStack(n *Node[T]) {
	// define the stack
	s := stack.NewStack[*Node[T]]()
	curr := t.Root
	for {
		if curr != nil {
			// push node to stack
			s.Push(curr)
			// go to left
			curr = curr.left
		} else {
			if s.IsEmpty() {
				break
			}
			// pop and print node
			n := s.PopAlt()
			fmt.Print(n.val)
			// go to right
			curr = curr.right
		}
	}
}

/*
Inorder traversal using Morris Traversal:
Initialize the current node as root.
While current is not null, check if it has a left child.
If there is no left child, print the current node and move to the right child of the current node.
Otherwise, find the rightmost node of the left subtree or the node whose right child is the current node:
If the right child is NULL, make current as the right child and move to the left child of current.
If the right child is the current node itself, print current node, make the right child NULL and move to the right child of the current node.
*/
func (t *Tree[T]) InorderTraverseMorris() {
	curr := t.Root
	for curr != nil {
		if curr.left == nil {
			// if there's no node to left
			// print the parent
			fmt.Print(curr.val)
			// go to right
			curr = curr.right
		} else {
			// left node exists
			tmp := curr.left
			for tmp.right != nil && tmp.right != curr {
				// iterate until right = null & right == curr
				// go to right
				tmp = tmp.right
			}
			if tmp.right == nil {
				// end of right reached
				tmp.right = curr
				curr = curr.left
			} else {
				// tmp.right == curr
				tmp.right = nil
				fmt.Print(curr.val)
				curr = curr.right
			}
		}
	}
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
Start with root node and push onto stack.
Repeat while the stack is not empty
POP the top element (curr) from the stack and process the node.
PUSH the right child of curr onto to stack.
PUSH the left child of curr onto to stack.
*/
func (t *Tree[T]) PreorderTraverseStack() {
	s := stack.NewStack[*Node[T]]()
	s.Push(t.Root)
	for !s.IsEmpty() {
		n := s.PopAlt()
		if n == nil {
			break
		}
		fmt.Println(n.val)
		if n.right != nil {
			s.Push(n.right)
		}
		if n.right != nil {
			s.Push(n.right)
		}
	}
}

/*
Morris traversal
1...If left child is null, print the current node data. Move to right child.
….Else, Make the right child of the inorder predecessor point to the current node. Two cases arise:
………a) The right child of the inorder predecessor already points to the current node. Set right child to NULL. Move to right child of current node.
………b) The right child is NULL. Set it to the current node. Print the current node’s data and move to left child of current node.
2...Iterate until the current node is not NULL.
*/
func (t *Tree[T]) PreorderTraverseMoris() {
	curr := t.Root
	for curr != nil {
		// If left child is null, print the current node data. Move to
		// right child.
		if curr.left == nil {
			fmt.Println(curr.val)
			curr = curr.right
		} else {
			curr = curr.left
			for curr.right != nil && curr.right != curr {
				curr = curr.right
			}
			// If the right child of inorder predecessor
			// already points to this node
			if curr.right == curr {
				curr.right = nil
				curr = curr.right
			} else {
				// If right child doesn't point to this node, then print
				// this node and make right child point to this node
				fmt.Println(curr.val)
				curr.right = curr
				curr = curr.left
			}
		}
	}
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

/*
1.1 Create an empty stack
2.1 Do following while root is not NULL
    a) Push root's right child and then root to stack.
    b) Set root as root's left child.
2.2 Pop an item from stack and set it as root.
    a) If the popped item has a right child and the right child 
       is at top of stack, then remove the right child from stack,
       push the root back and set root as root's right child.
    b) Else print root's data and set root as NULL.
2.3 Repeat steps 2.1 and 2.2 while stack is not empty.
*/
func (t *Tree[T]) PostorderTraverseStack() {
	s := stack.NewStack[*Node[T]]()
	curr := t.Root
	for curr != nil {
		if curr.right != nil {
			s.Push(curr.right)
		}
		curr = curr.left
	}
}

/*
 */
func (t *Tree[T]) PostorderTraverseMoris() {
	s := stack.NewStack[*Node[T]]()
	curr := t.Root
	for curr != nil {
		if curr.right != nil {
			s.Push(curr)
			curr = curr.left
		} else {
			prev := curr.right
			for prev.left != nil && prev.left != curr {
				prev = prev.left
			}
			if prev.left != nil {
				prev.left = curr
				s.Push(curr)
				curr = curr.right
			} else {
				curr = curr.left
				prev.left = nil
			}
		}
	}
	for !s.IsEmpty() {
		fmt.Println(s.PopAlt().val)
	}
}
