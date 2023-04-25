package algorithm

func BubbleSort(arr []int) []int {
	// loop to access each array element
	for i := 0; i < len(arr); i++ {
		// loop to compare array elements
		for j := 0; j < len(arr)-i-1; j++ {
			// this loop rearranges the array
			// swap arr[j] with arr[j+1] if arr[j] > arr[j+1]
			if arr[j] > arr[j+1] {
				// store current arr[j]
				tmp := arr[j]
				// set arr[j+1] val to arr[j]
				arr[j] = arr[j+1]
				// set previous arr[j] val to arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

func InsertionSort(arr []int) []int {
	// starts from 2nd element
	for i := 1; i < len(arr); i++ {
		// store arr[i] in key
		key := arr[i]
		// init 2nd pointer j = i - 1
		j := i - 1
		// compare key with all elements in sorted sub list
		// Compare key with each element on the left of it until an element smaller
		// than it is found.
		for j >= 0 && key < arr[j] {
			// Shift all the elements in the sorted sub-list
			// that is greater than the value to be sorted
			arr[j+1] = arr[j]
			j--
		}
		// Place key at after the element just smaller than it.
		// Insert the value
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
