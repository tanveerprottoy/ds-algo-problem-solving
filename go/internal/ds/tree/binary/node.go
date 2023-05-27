package binary

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	val   T
	left  *Node[T]
	right *Node[T]
}

func NewNode[T constraints.Ordered](val T, leftChild, rightChild *Node[T]) *Node[T] {
	return &Node[T]{val: val, left: leftChild, right: rightChild}
}
