package slice

func RemoveAt[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func AppendCopy[T any](src, dest []T) []T {
	return append(dest, src...)
}
