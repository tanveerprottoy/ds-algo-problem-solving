package binary

type Node[T any] struct {
	data       T
	leftChild  *Node[T]
	rightChild *Node[T]
}

/*
Create – create a tree in data structure.
Insert − Inserts data in a tree.
Search − Searches specific data  in a tree to check it is present or not.
Preorder Traversal – perform Traveling a tree in a pre-order manner in data structure .
In order Traversal – perform Traveling a tree in an in-order manner.
Post order Traversal –perform Traveling a tree in a post-order manner.
*/
type Tree[T any] struct {
	Node T
	Edge *T
}
