package general

import "math"

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

func ReverseInt(num int) int {
	res := 0
	for num > 0 {
		// need to multiply by ten
		// in res for each iteration
		// so that the number grows
		// from right to left
		// i.e. num = 124,
		// num % 10 = 4
		// 0 * 10 + 4
		// in 2nd iteration
		// num = 12,
		// num % 10 = 2
		// 4 * 10 + 2
		res = res*10 + num%10
		// 1st iteration: 124 / 10 = 12
		num /= 10
	}
	return res
}

func ReverseIntToArr(num int) []int {
	var res []int
	for num > 0 {
		res = append(res, num%10)
		num /= 10
	}
	return res
}

func IntFromArray(nums []int) int {
	num := 0
	// multiplier which will
	// make the target number
	// as will build the number from most
	// significant digit to least [left to right]
	// need make the multiplier from
	// the most significant digit
	// mul will be 10^(len(digits)-1)
	// i.e. [1,2,3] most significant digit is 1
	// and it has 3 digits, so need to make
	// multiplier as 10 ^ (3-1) = 100
	mul := int(math.Pow10(len(nums) - 1))
	for i := range nums {
		// extract each digit
		// need to multiply with
		// multiplier so that the
		// number builds up
		// i.e. [1,2,3]
		// only adding [1,2,3] will result 6
		// in case [1,2,3] the multiplier
		// will be 100
		// first num will be
		// 1 * 100
		num += nums[i] * mul
		// divide the mul by 10
		// for the next digit
		mul /= 10
	}
	return num
}
