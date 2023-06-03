package general

type Node[T any] struct {
	data   T
	parent *Node[T]
	left   *Node[T]
	right  *Node[T]
}

func NewNode[T any](data T, parent, left, right *Node[T]) *Node[T] {
	return &Node[T]{data: data, parent: parent, left: left, right: right}
}

type NodeAlt[T any] struct {
	data     T
	children []*NodeAlt[T]
}

func NewNodeAlt[T any](data T, children []*NodeAlt[T]) *NodeAlt[T] {
	return &NodeAlt[T]{data: data, children: children}
}
