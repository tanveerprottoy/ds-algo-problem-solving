package main

import (
	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/linkedlist/singly/generic"
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
	l3 := generic.NewNode(4, nil)
	l2 := generic.NewNode(3, l3)
	l1 := generic.NewNode(2, nil)
	l0 := generic.NewNode(1, l1)
	res := general.ZipLinkedLists(l0, l2)
	list := generic.NewLinkedList(res, nil)
	list.Traverse()
	res1 := general.ZipLinkedListsRecur(l0, l2)
	list = generic.NewLinkedList(res1, nil)
	list.Traverse()
}
