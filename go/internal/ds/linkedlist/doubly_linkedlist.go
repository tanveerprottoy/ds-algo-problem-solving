package linkedlist

type DoublyLinkedList[T any] struct {
	Val  T
	Nxt  *Node[T]
	Prev *Node[T]
}
