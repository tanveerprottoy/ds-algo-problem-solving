package main

import (
	"fmt"

	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/algorithm"
)

func main() {
	// fmt.Println(algorithm.BubbleSort([]int{3, 2, -2, 3}))
	// fmt.Println(algorithm.BubbleSort([]int{90, 111, 0, -2}))
	// fmt.Println(algorithm.BubbleSortOpt([]int{90, 111, 0, -2}))
	// fmt.Println(algorithm.BubbleSortDesc([]int{90, 111, 0, -2}))
	// fmt.Println(algorithm.SelectionSort([]int{111, 102, 90, 0, -2}))
	// fmt.Println(algorithm.Insertion([]int{111, 102, 90, 0, -2}))
	// fmt.Println(algorithm.Insertion([]int{111, 102, 90, 0, -2}))
	// arr := []int{111, 102, 90, 80, -2, 25, 0, -9}
	arr := []int{4, 3, 2, 1}
	// algorithm.MergeSort(arr, 0, len(arr)-1)
	// algorithm.MergeSortTest(arr, 0, len(arr)-1)
	algorithm.HeapSort(arr)
	fmt.Println(arr)
	// fmt.Println(algorithm.Bubble([]int{111, 102, 90, 0, -2}))
}
