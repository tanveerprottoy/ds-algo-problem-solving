package slice

import "reflect"

/*
The expression ~string means the set of all types whose underlying type is string.
This includes the type string itself as well as all types declared with definitions such
as type MyString string.
Here S must be a slice type whose element type can be any type.
*/
func RemoveAt[S ~[]E, E any](s S, index int) S {
	return append(s[:index], s[index+1:]...)
}

func AppendCopy[S ~[]E, E any](src, dest S) S {
	return append(dest, src...)
}

func Reverse[S ~[]E, E any](s S) S {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func ReverseSwap[S ~[]E, E any](s S) S {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	return s
}
