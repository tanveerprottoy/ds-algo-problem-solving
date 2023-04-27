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
			// Compare key with each element on the left of it
			// an element smaller than key is found.
			fmt.Println("inner loop pass: ", arr)
			// Shift all the elements in the sorted sub-list
			// that is greater than the value to be sorted
			arr[j+1] = arr[j]
			j--
			fmt.Println("inner loop pass after swap: ", arr)
		}
		// Place key at after the element just smaller than it.
		// Insert the value, j will be in the position where
		// item before is is smaller than k
		arr[j+1] = key
	}
	return arr
}

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			// To sort in descending order, change > to < in this line.
			// Select the minimum element in each loop.
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// put min at the correct position
		tmp := arr[i]
		arr[i] = arr[minIdx]
		arr[minIdx] = tmp
	}
	return arr
}

func Bubble(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		// outer
		for j := 0; j < len(arr)-i-1; j++ {
			// inner
			// this takes the largest value to the right end
			// with every completion pass
			// check adjacent values
			if arr[j] > arr[j+1] {
				// shift
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}
