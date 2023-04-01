package main

import (
	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/linkedlist"
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
	l3 := linkedlist.NewNode(4, nil)
	l2 := linkedlist.NewNode(3, l3)
	l1 := linkedlist.NewNode(2, nil)
	l0 := linkedlist.NewNode(1, l1)
	/* res := general.ZipLinkedLists(l0, l2)
	res.Traverse(res) */
	res1 := general.ZipLinkedListsRecur(l0, l2)
	res1.Traverse(res1)
}
