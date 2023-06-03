package binary

import (
	"fmt"

	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/stack"
	"github.com/tanveerprottoy/ds-algo-problem-solving/pkg/slice"
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
one type of DFS
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
2 Push curr onto stack if curr is not nil.
3 Move to left of curr and repeat step 2.
4 If curr is nil and stack is not empty, then Pop element from stack and set as curr.
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
While current is not nil, check if it has a left child.
If there is no left child, print the current node and move to the right child of the current node.
Otherwise, find the rightmost node of the left subtree or the node whose right child is the current node:
If the right child is nil, make current as the right child and move to the left child of current.
If the right child is the current node itself, print current node, make the right child nil and move to the right child of the current node.
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
				// iterate until right = nil & right == curr
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
one type of DFS
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
1...If left child is nil, print the current node data. Move to right child.
….Else, Make the right child of the inorder predecessor point to the current node. Two cases arise:
………a) The right child of the inorder predecessor already points to the current node. Set right child to nil. Move to right child of current node.
………b) The right child is nil. Set it to the current node. Print the current node’s data and move to left child of current node.
2...Iterate until the current node is not nil.
*/
func (t *Tree[T]) PreorderTraverseMoris() {
	curr := t.Root
	for curr != nil {
		// If left child is nil, print the current node data. Move to
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
one type of DFS
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
2.1 Do following while root is not nil

	a) Push root's right child and then root to stack.
	b) Set root as root's left child.

2.2 Pop an item from stack and set it as root.

	a) If the popped item has a right child and the right child
	   is at top of stack, then remove the right child from stack,
	   push the root back and set root as root's right child.
	b) Else print root's data and set root as nil.

2.3 Repeat steps 2.1 and 2.2 while stack is not empty.
*/
func (t *Tree[T]) PostorderTraverse1Stack() {
	// Check for empty tree
	if t.Root == nil {
		return
	}
	s := stack.NewStack[*Node[T]]()
	var data []*Node[T]
	var prev *Node[T]
	s.Push(t.Root)
	for !s.IsEmpty() {
		current := s.Peek()
		/* go down the tree in search of a leaf an if so
		   process it and pop stack otherwise move down */
		if prev == nil || prev.left == current || prev.right == current {
			if current.left != nil {
				s.Push(current.left)
			} else if current.right != nil {
				s.Push(current.right)
			} else {
				s.PopAlt()
				data = append(data, current)
			}
			/* go up the tree from left node, if the
			   child is right push it onto stack otherwise
			   process parent and pop stack */
		} else if current.left == prev {
			if current.right != nil {
				s.Push(current.right)
			} else {
				s.PopAlt()
				data = append(data, current)
			}
			/* go up the tree from right node and after
			   coming back from right node process parent
			   and pop stack */
		} else if current.right == prev {
			s.PopAlt()
			data = append(data, current)
		}
		prev = current
	}
	for _, e := range data {
		fmt.Println(e.val)
	}
}

/*
 1. Push root to first stack.
 2. Loop while first stack is not empty
    2.1 Pop a node from first stack and push it to second stack
    2.2 Push left and right children of the popped node to first stack
 3. Print contents of second stack
*/
func (t *Tree[T]) PostorderTraverse2Stacks() {
	s1 := stack.NewStack[*Node[T]]()
	s2 := stack.NewStack[*Node[T]]()
	if t.Root == nil {
		return
	}
	// push root to first stack
	s1.Push(t.Root)
	// Run while first stack is not empty
	for !s1.IsEmpty() {
		// Pop an item from s1 and push it to s2
		temp := s1.PopAlt()
		s2.Push(temp)
		// Push left and right children of
		// removed item to s1
		if temp.left != nil {
			s1.Push(temp.left)
		}
		if temp.right != nil {
			s1.Push(temp.right)
		}
	}
	// Print all elements of second stack
	for !s2.IsEmpty() {
		temp := s2.PopAlt()
		fmt.Println(temp.val)
	}
}

/*
Create a slice and Initialize current as root
While current is not nil
If the current does not have a right child
push the current key in slice

	Go to the left, i.e., current = current->left

else
Find leftmost node in current right subtree OR node whose left child == current.
if current  does not have a left child
push the current node in the slice
Make current as of the left child of that leftmost node
Go to this right child, i.e., current = current->right
else
Found left child == current
Update the left child as nil of that node whose left child is current
Go to the left, i.e. current = current->left

In the last, we reverse the slice and print it, since slice is used to store the output, the space complexity of this algorithm would be O(N).
*/
func (t *Tree[T]) PostorderTraverseMoris() {
	var data []*Node[T]
	current := t.Root
	for current != nil {
		// If right child is nil,
		// put the current node data
		// in res. Move to left child.
		if current.right == nil {
			data = append(data, current)
			current = current.left
		} else {
			predecessor := current.right
			for predecessor.left != nil && predecessor.left != current {
				predecessor = predecessor.left
			}
			// If left child doesn't point
			// to this node, then put in res
			// this node and make left
			// child point to this node
			if predecessor.left == nil {
				data = append(data, current)
				predecessor.left = current
				current = current.right
			} else {
				// If the left child of inorder predecessor
				// already points to this node
				predecessor.left = nil
				current = current.left
			}
		}
	}
	// reverse the res
	slice.Reverse(data)
	for _, e := range data {
		fmt.Println(e.val)
	}
}

/*
Print the level order traversal of the tree using recursive function to traverse all nodes of a level.
Find height of tree and run depth first search and maintain current height, print nodes for every 
height from root and for 1 to height and match if the current height is equal to height of the 
iteration then print node’s data.
*/
func (t *Tree[T]) LevelOrderTraversal() {

}
