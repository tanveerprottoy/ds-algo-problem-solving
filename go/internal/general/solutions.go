package general

func BubbleSort(arr []int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := 0; j < i-1; j++ {

		}
	}
}

func Reverse(arr []int) []int {
	// 1st assign, 2nd condition, 3rd increment, decrement
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func Factorial(n int32) int32 {
	if n < 2 {
		return n
	}
	return n * Factorial(n-1)
}

// minimum-number-of-distinct-elements-after-removing-m-items
func MinimumNumberOfDistinctElements(arr []int) {
}
