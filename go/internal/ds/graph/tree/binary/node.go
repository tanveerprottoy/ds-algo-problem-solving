package binary

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	val   T
	left  *Node[T]
	right *Node[T]
}

func NewNode[T constraints.Ordered](val T, left, right *Node[T]) *Node[T] {
	return &Node[T]{val: val, left: left, right: right}
}
