package heap

// Get parent index
func Parent(i int) int {
	return (i - 1) / 2
}

// Get the left child index
func Left(i int) int {
	return 2*i + 1
}

// Get the right child index
func Right(i int) int {
	return 2*i + 2
}
