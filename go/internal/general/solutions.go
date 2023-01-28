package general

func Reverse(arr []int) []int {
	// 1st assign, 2nd condition, 3rd increment, decrement
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
