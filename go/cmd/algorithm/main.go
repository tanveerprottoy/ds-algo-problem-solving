package main

import (
	"fmt"

	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/algorithm"
)

func main() {
	// fmt.Println(algorithm.BubbleSort([]int{3, 2, -2, 3}))
	// fmt.Println(algorithm.BubbleSort([]int{90, 111, 0, -2}))
	// fmt.Println(algorithm.BubbleSortOpt([]int{90, 111, 0, -2}))
	//fmt.Println(algorithm.Bubble([]int{90, 111, 0, -2}))
	fmt.Println(algorithm.Insertion([]int{111, 102, 90, 0, -2}))
}
