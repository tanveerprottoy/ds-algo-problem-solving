package ds

type BasicLinkedList struct {
	Val int
	Nxt *BasicLinkedList
}

type LinkedList[T any] struct {
	Val T
	Nxt *LinkedList[T]
}

func (*l BasicLinkedList) Traverse() {

}