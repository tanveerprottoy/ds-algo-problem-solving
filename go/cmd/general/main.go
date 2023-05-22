package main

import (
	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/linkedlist/singly"
	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/general"
)

func main() {
	/* fmt.Printf("%d\n", general.Factorial(4))
	fmt.Printf("%d\n", general.Factorial(2))
	fmt.Printf("%d\n", general.Factorial(0))
	fmt.Printf("%d\n", general.Factorial(1)) */
	// fmt.Println(general.ReverseIntToArr(124))
	/* fmt.Println(general.ReverseInt(124))
	fmt.Println(general.IntFromArray([]int{1, 2, 4})) */
	/* fmt.Println(general.FibonacciInefficient(5))
	fmt.Println(general.FibonacciInefficient(8))
	fmt.Println(general.FibonacciInefficient(50)) */
	// fmt.Println(general.SumNonNegativeNums(5))
	// fmt.Println(general.SumNonNegativeNums(10000000))
	l3 := singly.NewNode(4, nil)
	l2 := singly.NewNode(3, l3)
	l1 := singly.NewNode(2, nil)
	l0 := singly.NewNode(1, l1)
	res := general.ZipLinkedLists(l0, l2)
	list := singly.NewLinkedList(res)
	list.Traverse()
	res1 := general.ZipLinkedListsRecur(l0, l2)
	list = singly.NewLinkedList(res1)
	list.Traverse()
}
