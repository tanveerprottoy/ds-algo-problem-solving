package algorithm

import (
	"fmt"
)

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
Merge Sort is one of the most popular sorting algorithms that is based
on the principle of Divide and Conquer Algorithm.
Here, a problem is divided into multiple sub-problems. Each sub-problem
is solved individually. Finally, sub-problems are combined to form the final solution.

In simple terms, we can say that the process of merge sort is to divide the array into two halves,
sort each half, and then merge the sorted halves back together. This process is repeated until
the entire array is sorted.

Have we reached the end of any of the arrays?

	No:
	    Compare current elements of both arrays
	    Copy smaller element into sorted array
	    Move pointer of element containing smaller element
	Yes:
	    Copy all remaining elements of non-empty array

Suppose we had to sort an array A. A subproblem would be to sort a sub-section of this array starting
at index p and ending at index r, denoted as A[p..r].

# Divide

If q is the half-way point between p and r, then we can split the subarray A[p..r] into two
arrays A[p..q] and A[q+1, r].

# Conquer

In the conquer step, we try to sort both the subarrays A[p..q] and A[q+1, r]. If we haven't yet
reached the base case, we again divide both these subarrays and try to sort them.

# Combine

When the conquer step reaches the base step and we get two sorted subarrays A[p..q] and A[q+1, r]
for array A[p..r], we combine the results by creating a sorted array A[p..r] from two sorted
subarrays A[p..q] and A[q+1, r].

The merge function works as follows:

Create copies of the subarrays L <- A[p..q] and M <- A[q+1..r].
Create three pointers i, j and k
i maintains current index of L, starting at 1
j maintains current index of M, starting at 1
k maintains the current index of A[p..q], starting at p.
Until we reach the end of either L or M, pick the larger among the elements from L and M and
place them in the correct position at A[p..q]
When we run out of elements in either L or M, pick up the remaining elements and put in A[p..q]
*/
func Merge(arr []int, left, mid, right int) {
	// calculate sizes of two subarrays to be merged
	// for 1st its end is mid and start is left
	// for 2nd its end is right and start is mid
	fmt.Print("\n")
	fmt.Println("Merge")
	//fmt.Println("arr: ", arr)
	fmt.Println("left: ", left)
	fmt.Println("mid: ", mid)
	fmt.Println("right: ", right)
	fmt.Print("\n")
	lenLeft := mid - left + 1
	lenRight := right - mid
	// init temp arrays
	leftArr := make([]int, lenLeft)
	rightArr := make([]int, lenRight)
	// fill temp arrays with values
	for i := 0; i < lenLeft; i++ {
		leftArr[i] = arr[left+i]
	}
	for i := 0; i < lenRight; i++ {
		rightArr[i] = arr[mid+1+i]
	}
	fmt.Println("leftArr: ", leftArr)
	fmt.Println("rightArr: ", rightArr)
	// initial index of 1st sub-array
	i := 0
	// initial index of 2nd sub-array
	j := 0
	// initial index of merged sub-array
	k := left
	// until reached either end of either L or M, pick larger among
	// elements leftArr and rightArr and place them in the correct position on a
	for i < lenLeft && j < lenRight {
		// check if leftArr[i] < rightArr[j]
		if leftArr[i] <= rightArr[j] {
			// set arr[k] = leftArr[i]
			// that way smaller value is placed
			// at the correct position
			arr[k] = leftArr[i]
			// do increment 1st sub arr index
			// when merging array item was placed
			// from it
			i++
		} else {
			// otherwise
			// set arr[k] = rightArr[j]
			// as rightArr[j] is the smaller value
			arr[k] = rightArr[j]
			// do increment 2nd sub arr index
			// when merging array item was placed
			// from it
			j++
		}
		// for each iteration increase the merged arr index
		k++
	}
	// the loop might stop as i < leng0 && j < leng1
	// with one of them finishing to it's target value
	// for that need to check if if any item remains on the
	// leftArr & rightArr respectively
	// check leftArr
	for i < lenLeft {
		arr[k] = leftArr[i]
		i++
		k++
	}
	// check rightArr
	for j < lenRight {
		arr[k] = rightArr[j]
		j++
		k++
	}
	fmt.Println("arr: ", arr)
	fmt.Print("\n")
}

func MergeSort(arr []int, left, right int) {
	// []int{4, 3, 2, 1}
	/* fmt.Print("\n")
	fmt.Println("MergeSort")
	fmt.Println("arr: ", arr)
	fmt.Println("left: ", left)
	fmt.Println("right: ", right)
	fmt.Print("\n") */
	if left >= right {
		return
	}
	mid := (left + right) / 2
	// fmt.Println("mid: ", mid)
	// left half
	MergeSort(arr, left, mid)
	// right half
	MergeSort(arr, mid+1, right)
	Merge(arr, left, mid, right)
}

func MergeAlt(leftArr, rightArr, arr []int) {
	lenLeft := len(leftArr)
	lenRight := len(rightArr)
	fmt.Println("leftArr: ", leftArr)
	fmt.Println("rightArr: ", rightArr)
	// initial index of 1st sub-array
	i := 0
	// initial index of 2nd sub-array
	j := 0
	// initial index of merged sub-array
	k := 0
	// until reached either end of either L or M, pick larger among
	// elements leftArr and rightArr and place them in the correct position on a
	for i < lenLeft && j < lenRight {
		// check if leftArr[i] < rightArr[j]
		if leftArr[i] <= rightArr[j] {
			// set arr[k] = leftArr[i]
			// that way smaller value is placed
			// at the correct position
			arr[k] = leftArr[i]
			// do increment 1st sub arr index
			// when merging array item was placed
			// from it
			i++
		} else {
			// otherwise
			// set arr[k] = rightArr[j]
			// as rightArr[j] is the smaller value
			arr[k] = rightArr[j]
			// do increment 2nd sub arr index
			// when merging array item was placed
			// from it
			j++
		}
		// for each iteration increase the merged arr index
		k++
	}
	// the loop might stop as i < leng0 && j < leng1
	// with one of them finishing to it's target value
	// for that need to check if if any item remains on the
	// leftArr & rightArr respectively
	// check leftArr
	for i < lenLeft {
		arr[k] = leftArr[i]
		i++
		k++
	}
	// check rightArr
	for j < lenRight {
		arr[k] = rightArr[j]
		j++
		k++
	}
	fmt.Println("arr: ", arr)
	fmt.Print("\n")
}

func MergeSortAlt(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	mid := length / 2
	leftArr := make([]int, mid)
	rightArr := make([]int, length-mid)
	// fill temp arrays with values
	for i := 0; i < mid; i++ {
		leftArr[i] = arr[i]
	}
	for i := 0; i < length-mid; i++ {
		rightArr[i] = arr[i+mid]
	}
	// left half
	MergeSortAlt(leftArr)
	// right half
	MergeSortAlt(rightArr)
	MergeAlt(leftArr, rightArr, arr)
}

func MergeTest(arr []int, left, mid, right int) {
	lenLeft := mid - left + 1
	lenRight := right - mid
	leftArr := make([]int, lenLeft)
	rightArr := make([]int, lenRight)
	for i := 0; i < lenLeft; i++ {
		leftArr[i] = arr[i+left]
	}
	for i := 0; i < lenRight; i++ {
		rightArr[i] = arr[i+mid+1]
	}
	i, j, k := 0, 0, left
	for i < lenLeft && j < lenRight {
		if leftArr[i] < rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = leftArr[j]
			j++
		}
		k++
	}
	for i < lenLeft {
		arr[k] = leftArr[i]
		i++
		k++
	}
	for j < lenRight {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

func MergeSortTest(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := len(arr) / 2
	MergeSortTest(arr, left, mid)
	MergeSortTest(arr, mid+1, right)
	MergeTest(arr, left, mid, right)
}

func MergeTestAlt(leftArr, rightArr, arr []int) {
	lenLeft := len(leftArr)
	lenRight := len(rightArr)
	i, j, k := 0, 0, 0
	for i < lenLeft && j < lenRight {
		if leftArr[i] < rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}
	for i < lenLeft {
		arr[k] = leftArr[i]
		i++
		k++
	}
	for j < lenRight {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

func MergeSortTestAlt(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	mid := length / 2
	leftArr := make([]int, mid)
	rightArr := make([]int, length-mid)
	for i := 0; i < mid; i++ {
		leftArr[i] = arr[i]
	}
	for i := 0; i < length-mid; i++ {
		rightArr[i] = arr[i+mid]
	}
	MergeSortTestAlt(leftArr)
	MergeSortTestAlt(rightArr)
	MergeTestAlt(leftArr, rightArr, arr)
}

/*
	Quicksort is a sorting algorithm based on the divide and conquer approach where

An array is divided into subarrays by selecting a pivot element (element selected from the array).

While dividing the array, the pivot element should be positioned in such a way that elements less than pivot are kept on the left side and elements greater than pivot are on the right side of the pivot.
The left and right subarrays are also divided using the same approach. This process continues until each subarray contains a single element.
At this point, elements are already sorted. Finally, elements are combined to form a sorted array.
*/
func Partition(arr []int, left, right int) int {
	// store the pivot
	pivot := arr[right]
	// smaller than pivot item found pointer
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			// if smaller element found than pivot
			// increment i first
			i++
			// swap it with arr[i]
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp
		}
	}
	// swap pivot with arr[i+1]
	tmp := arr[i+1]
	arr[i+1] = arr[right]
	arr[right] = tmp
	return i + 1
}

func QuickSort(arr []int, left, right int) {
	for left >= right {
		return
	}
	pivotIdx := Partition(arr, left, right)
	QuickSort(arr, left, pivotIdx-1)
	QuickSort(arr, pivotIdx+1, right)
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

If the index of any element in the array is i, the element in the index 2i+1 will become the 
left child and element in 2i+2 index will become the right child. Also, the parent of any 
element at index i is given by the lower bound of (i-1)/2.

Left child of 1 (index 0)
= element in (2*0+1) index 
= element in 1 index 
= 12


Right child of 1
= element in (2*0+2) index
= element in 2 index 
= 9

Similarly,
Left child of 12 (index 1)
= element in (2*1+1) index
= element in 3 index
= 5

Right child of 12
= element in (2*1+2) index
= element in 4 index
= 6
Let us also confirm that the rules hold for finding parent of any node

Parent of 9 (position 2) 
= (2-1)/2 
= ½ 
= 0.5
~ 0 index 
= 1

Parent of 12 (position 1) 
= (1-1)/2 
= 0 index 
= 1
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
		swapped := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}

func Insertion(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			tmp := arr[j]
			arr[j] = arr[j+1]
			arr[j+1] = tmp
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func Selection(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i - 1; j >= 0; j-- {
			if arr[j] < arr[j+1] {
				minIdx = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[minIdx]
		arr[minIdx] = tmp
	}
	return arr
}
