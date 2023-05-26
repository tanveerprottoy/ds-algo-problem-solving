package nonbinary

type Node[T any] struct {
	data         T
	parent       *Node[T]
	leftChild    *Node[T]
	rightSibling *Node[T]
}

func NewNode[T any](data T, parent, leftChild, rightSibling *Node[T]) *Node[T] {
	return &Node[T]{data: data, parent: parent, leftChild: leftChild, rightSibling: rightSibling}
}

type NodeAlt[T any] struct {
	data     T
	children []*NodeAlt[T]
}

func NewNodeAlt[T any](data T, children []*NodeAlt[T]) *NodeAlt[T] {
	return &NodeAlt[T]{data: data, children: children}
}
