package algorithm

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i-1; j++ {
			// this loop rearranges
			// the array
			// swap arr[j] with arr[j+1] if arr[j] > arr[j+1]
			if arr[j] > arr[j+1] {
				// store arr[j]
				tmp := arr[j]
				// set arr[j] to arr[j+1]
				arr[j] = arr[j+1]
				// set previous arr[j] to arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

func InsertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i-1; j++ {
			// this loop rearranges
			// the array
			// get arr[j]
			c := arr[j]
			// get next arr[j+1]
			n := arr[j+1]
			// swap c with n if c > n
			if c > n {
				// set next val to arr[j]
				arr[j] = n
				// set current val to arr[j+1]
				arr[j+1] = c
			}
		}
	}
	return arr
}

func SelectionSort(arr []int) []int {
	return int[]]{}
}