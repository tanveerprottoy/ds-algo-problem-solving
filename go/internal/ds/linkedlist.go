package ds

import "fmt"

type BasicLinkedList struct {
	Val int
	Nxt *BasicLinkedList
}

func (l *BasicLinkedList) Traverse() {
	tmp := l
	for tmp != nil {
		fmt.Println(tmp.Val)
		fmt.Println(tmp.Nxt)
		tmp = tmp.Nxt
	}
}

type LinkedList[T any] struct {
	Val T
	Nxt *LinkedList[T]
}
