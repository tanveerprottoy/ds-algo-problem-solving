package algorithm

func BubbleSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l; i++ {
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
