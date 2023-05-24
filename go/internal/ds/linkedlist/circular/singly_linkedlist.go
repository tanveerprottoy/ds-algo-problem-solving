package circular

import (
	"fmt"
)

type SinglyNode struct {
	Val  any
	Next *SinglyNode
}

func NewSinglyNode(val any, Next *SinglyNode) *SinglyNode {
	return &SinglyNode{Val: val, Next: Next}
}

type SinglyLinkedList struct {
	Head   *SinglyNode
	Tail   *SinglyNode
	Length int
}

func NewSinglyLinkedList(head, tail *SinglyNode) *SinglyLinkedList {
	return &SinglyLinkedList{Head: head, Tail: tail}
}

func (l *SinglyLinkedList) Traverse() {
	curr := l.Head
	for curr != nil {
		fmt.Println("val: ", curr.Val)
		fmt.Println("Next: ", curr.Next)
		curr = curr.Next
	}
}

func (l *SinglyLinkedList) TraverseRecur(node *SinglyNode) {
	if node.Next == nil {
		return
	}
	fmt.Println("val: ", node.Val)
	fmt.Println("Next: ", node.Next)
	l.TraverseRecur(node.Next)
}

func (l *SinglyLinkedList) Size() int {
	if l.Head == nil {
		return 0
	}
	i := 1
	curr := l.Head
	for curr != nil {
		i++
		curr = curr.Next
	}
	return i
}

func (l *SinglyLinkedList) InsertAtHead(v any) {
	nxt := l.Head.Next
	n := NewSinglyNode(v, nxt)
	l.Head = n
	l.Length++
}

func (l *SinglyLinkedList) InsertAtTail(v any) {
	curr := l.Head
	for curr != nil {
		curr = curr.Next
	}
	n := NewSinglyNode(v, nil)
	curr.Next = n
	l.Tail = n
	l.Length++
}

func (l *SinglyLinkedList) InsertAtPosition(v any, pos int) {
	i := 0
	curr := l.Head
	for i < pos && curr != nil {
		if i == pos-1 {
			nxt := curr.Next
			n := NewSinglyNode(v, nxt)
			curr.Next = n
			l.Length++
			return
		}
		i++
		curr = curr.Next
	}
}

func (l *SinglyLinkedList) InsertAtMiddle(v any) {
	mid := l.Size() / 2
	i := 0
	curr := l.Head
	for curr != nil {
		if i == mid-1 {
			n := NewSinglyNode(v, curr.Next)
			n.Next = curr
			l.Length++
			return
		}
		curr = curr.Next
		i++
	}
}

func (l *SinglyLinkedList) Find(v any) *SinglyNode {
	curr := l.Head
	for curr != nil {
		switch v.(type) {
		case int:
			fmt.Print(v == curr.Val.(int))
			if v == curr.Val.(int) {
				return curr
			}
		}
		// val, ok := v.(int)
		curr = curr.Next
	}
	return nil
}

func (l *SinglyLinkedList) DeleteAtHead() *SinglyNode {
	curr := l.Head
	l.Head = l.Head.Next
	l.Length--
	return curr
}

func (l *SinglyLinkedList) DeleteAtTail() *SinglyNode {
	if l.Head == nil {
		return nil
	}
	var del *SinglyNode
	// corener case for only 1 item
	if l.Head.Next == nil {
		del := l.Head
		l.Head = nil
		l.Length--
		return del
	}
	curr := l.Head
	for curr != nil {
		curr = curr.Next
	}
	del = curr
	curr = nil
	l.Length--
	return del
}

func (l *SinglyLinkedList) Delete(v any) *SinglyNode {
	if l.Head == nil {
		return nil
	}
	if l.Head.Next == nil {
		tmp := l.Head
		l.Head = nil
		l.Length--
		return tmp
	}
	// slow will be the previous
	// node of fast
	fast := l.Head.Next
	slow := l.Head
	// check for conrner case
	// when head is the target value
	/* if v == slow.Val.(int) {
		slow = slow.Next
	} */
	for fast != nil && slow != nil {
		switch v.(type) {
		case int:
			{
				if v == fast.Val.(int) {
					slow.Next = fast.Next
					l.Length--
					return fast
				} else if v == slow.Val.(int) {
					slow = slow.Next
					l.Length--
					return slow
				}
			}
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return nil
}

func (l *SinglyLinkedList) insertionSortHelper(head, newNode *SinglyNode) *SinglyNode {
	tmp := NewSinglyNode(0, nil)
	curr := tmp
	tmp.Next = head
	// checking the current node with each node of the linked
	// list if it is smaller then we are swapping.
	for curr.Next != nil && curr.Next.Val.(int) < newNode.Val.(int) {
		curr = curr.Next
	}
	newNode.Next = curr.Next
	curr.Next = newNode
	return tmp.Next
}

func (l *SinglyLinkedList) InsertionSort() *SinglyNode {
	// insertion sort
	var res *SinglyNode
	curr := l.Head
	for curr != nil {
		res = l.insertionSortHelper(res, curr)
		curr = curr.Next
	}
	return res
}
