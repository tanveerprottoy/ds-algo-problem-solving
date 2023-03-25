package linkedlist

type Node[T any] struct {
	Val T
	Nxt *Node[T]
}