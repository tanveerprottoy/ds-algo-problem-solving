package algorithm

import "fmt"

/*
Bubble sort is a sorting algorithm that compares two adjacent
elements and swaps them until they are in the intended order.
Just like the movement of air bubbles in the water that rise
up to the surface, each element of the array move to the end
in each iteration. Therefore, it is called a bubble sort.
*/
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		// outer loop is for access each array element
		// outer loop handles the main pass
		fmt.Println("outer loop pass: ", arr)
		for j := 0; j < len(arr)-i-1; j++ {
			// j < len(arr)-i-1 as arr[j+1]
			// is accessed need to loop till
			// largest index - 1
			// loop to compare array elements
			// this loop rearranges the array
			// inner loop is shrinked by 1 position
			// as len(arr)-i-1, when i increases with
			// the outer loop
			// this is because with each outer loop pass
			// the largest number is sorted to the right
			// in ascending sort, so leaving the last most
			// index with each pass
			// swap arr[j] with arr[j+1] if arr[j] > arr[j+1]
			// change > to < to sort in descending order
			// ex: for descending if arr[j] < arr[j+1]
			fmt.Println("inner loop idx: ", j)
			fmt.Println("inner loop pass: ", arr)
			if arr[j] > arr[j+1] {
				// store current arr[j]
				tmp := arr[j]
				// set arr[j+1] val to arr[j]
				arr[j] = arr[j+1]
				// set previous arr[j] val to arr[j+1]
				arr[j+1] = tmp
			}
			fmt.Println("inner loop pass after swap: ", arr)
		}
	}
	return arr
}

func BubbleSortOpt(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		// outer loop is for access each array element
		// outer loop handles the main pass
		swapped := false
		fmt.Println("outer loop pass: ", arr)
		for j := 0; j < len(arr)-i-1; j++ {
			// loop to compare array elements
			// this loop rearranges the array
			// inner loop is shrinked by 1 position
			// as len(arr)-i-1, when i increases with
			// the outer loop
			// this is because with each outer loop pass
			// the largest number is sorted to the right
			// in ascending sort, so leaving the last most
			// index with each pass
			// swap arr[j] with arr[j+1] if arr[j] > arr[j+1]
			// change > to < to sort in descending order
			// ex: for descending if arr[j] < arr[j+1]
			fmt.Println("inner loop pass: ", arr)
			if arr[j] > arr[j+1] {
				// store current arr[j]
				tmp := arr[j]
				// set arr[j+1] val to arr[j]
				arr[j] = arr[j+1]
				// set previous arr[j] val to arr[j+1]
				arr[j+1] = tmp
				swapped = true
			}
		}
		fmt.Println("inner loop pass after swap: ", arr)
		if !swapped {
			// no adjacent elements were swapped by inner loop
			// break the outer loop
			fmt.Println("!swapped: ", swapped)
			break
		}
	}
	return arr
}

func BubbleSortDesc(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] < arr[j+1] {
				// shift
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

/*
Insertion sort is a sorting algorithm that places an unsorted
element at its suitable place in each iteration.
Insertion sort works similarly as we sort cards in our
hand in a card game.
We assume that the first card is already sorted then,
we select an unsorted card. If the unsorted card is
greater than the card in hand, it is placed on the
right otherwise, to the left. In the same way, other
unsorted cards are taken and put in their right place.
*/
func InsertionSort(arr []int) []int {
	// starts from 2nd element
	for i := 1; i < len(arr); i++ {
		// store arr[i] in key
		key := arr[i]
		// init 2nd pointer j = i - 1
		// this will compare key with values left side of it
		j := i - 1
		// compare key with all elements in sorted sub list
		// sorted sub list is every item left side of the key
		fmt.Println("outer loop pass: ", arr)
		for j >= 0 && key < arr[j] {
			// Shift all the elements in the sorted sub-list
			// that is greater than the value to be sorted
			// to the right
			// so for the first iteration of this inner loop
			// arr[j+1] = arr[j] swaps item to where current key
			// is, as j := i - 1
			// ex: key := arr[1]
			// j := 1 - 1
			// arr[0+1] = arr[0]
			// it places swapped value into index 1
			// where the key was
			// Compare key with each element on the left of it
			// an element smaller than key is found.
			fmt.Println("inner loop pass: ", arr)
			arr[j+1] = arr[j]
			j--
			fmt.Println("inner loop pass after swap: ", arr)
		}
		// Place key at after the element just smaller than it.
		// Insert the value, j will be in the position where
		// item before is is smaller than k
		// finally the key will be set in the right place
		// where the previous value is just smaller than the key
		// ex: arr[-1+1] = key
		arr[j+1] = key
	}
	return arr
}

func InsertionSortDesc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && key > arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

/*
Selection sort is a sorting algorithm that selects the smallest element from
an unsorted list in each iteration and places that element at the beginning of
the unsorted list.
*/
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		// set initial min index as current index
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			// start inner loop from i + 1
			// this loop finds the smallest value starting from
			// the adjacent right of the i/minIdx, j := i + 1
			// find if current arr[j] is smaller than arr[minIdx]
			// To sort in descending order, change > to < in this line.
			if arr[j] < arr[minIdx] {
				// if arr[j] < arr[minIdx] then set minIdx = j
				minIdx = j
			}
		}
		// swap arr[i] with arr[minIdx]
		tmp := arr[i]
		arr[i] = arr[minIdx]
		arr[minIdx] = tmp
	}
	return arr
}

func SelectionSortDesc(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		// set initial min index as current index
		maxIdx := i
		for j := i + 1; j < len(arr); j++ {
			// start inner loop from i + 1
			// this loop finds the largest value starting from
			// the adjacent right of the i/maxIdx, j := i + 1
			// find if current arr[j] is larger than arr[maxIdx]
			if arr[maxIdx] > arr[j] {
				// if arr[maxIdx] > arr[j] then set maxIdx = j
				maxIdx = j
			}
		}
		// swap arr[i] with arr[maxIdx]
		tmp := arr[i]
		arr[i] = arr[maxIdx]
		arr[maxIdx] = tmp
	}
	return arr
}

/*
Merge Sort is one of the most popular sorting algorithms that is based on the principle
of Divide and Conquer Algorithm.
Here, a problem is divided into multiple sub-problems. Each sub-problem is solved individually.
Finally, sub-problems are combined to form the final solution.

tips: 
Have we reached the end of any of the arrays?
    No:
        Compare current elements of both arrays 
        Copy smaller element into sorted array
        Move pointer of element containing smaller element
    Yes:
        Copy all remaining elements of non-empty array
*/
func MergeSort(arr []int) []int {
	return arr
}

func CountingSort(arr []int) []int {
	return arr
}

/*
Heapsort is an efficient sorting algorithm based on the
use of max/min heaps. A heap is a tree-based data structure
that satisfies the heap property – that is for a max heap,
the key of any node is less than or equal to the key
of its parent (if it has a parent).
*/
func HeapSort(arr []int) []int {
	return arr
}

func Bubble(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

func BubbleOpt(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		isSwapped := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
				isSwapped = true
			}
		}
		if !isSwapped {
			break
		}
	}
	return arr
}

func Insertion(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func Selection(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[minIdx]
		arr[minIdx] = tmp
	}
	return arr
}
