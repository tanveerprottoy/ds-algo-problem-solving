package algorithm

func BubbleSort(arr []int) []int {
	// loop to access each array element
	for i := 0; i < len(arr); i++ {
		// loop to compare array elements
		for j := 0; j < i-1; j++ {
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
	return []int{}
}

func SelectionSort(arr []int) []int {
	return []int{}
}
