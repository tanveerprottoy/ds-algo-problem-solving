package general

import (
	"math"

	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/linkedlist/singly"
)

func Reverse(arr []int) []int {
	// 1st assign, 2nd condition, 3rd increment, decrement
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func FactorialInefficient(n int32) int32 {
	if n < 2 {
		// stop here, as no need
		// to calculate <=1
		return n
	}
	// this algo decreases the
	// input by 1 and multiply
	// with n to form the factorial
	return n * FactorialInefficient(n-1)
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

func FibonacciInefficient(n int64) int64 {
	// formula: Fn = Fn-1+Fn-2
	if n <= 1 {
		// stop at 0 and 1, as seq
		// starts from 0, 1
		return n
	}
	// this algo calculates
	// fib seq to the nth
	// it decreases the
	// input by 1 and 2 respectively
	// then add them up to calcualate
	// the seq
	// i.e. f(5) = f(5 - 1) + f(5 - 2)
	//			 => f(4 - 1) + f(4 - 2)...
	return FibonacciInefficient(n-1) + FibonacciInefficient(n-2)
}

/*
		given a number n sum all nonnegative numbers upto n
		# tips for recursive problem
	 	1. first think of the simplest possible input, for this
		prob simplest input is n = 0, the simplest possible input
		most often play as the base case for the recursive algo
		2. visaulize the inputs and outputs, write down and form the tree
		for the recursive algo
		3. compare hard cases for the larger input with the simpler cases
		for smaller input, find the pattern in them, in recursive algo
		often the computation builds upon from smaller to larger blocks
		4. figure out and generalize the pattern, for this prob
		will have to calculate sum(n - 1) then add n to it
		i.e. n + sum(n - 1)
*/
func SumNonNegativeNums(n uint32) uint32 {
	// base case
	if n == 0 {
		return n
	}
	return n + SumNonNegativeNums(n-1)
}

func ZipLinkedLists(l1, l2 *singly.Node[int]) *singly.Node[int] {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// result traversal, mutating the l1,
	// result will be l1
	tmp0 := l1
	// l1 traversal, starting point must be
	// next as tmp0 is assigned to l1
	tmp1 := l1.Next
	// l2 traversal
	tmp2 := l2
	// use toggler to toggle between two lists
	// if even read tmp2, tmp1 when odd, as tmp0
	// starts from l1, need to read from l2 first
	toggler := 0
	// traverse till one of the lists is empty
	for tmp0 != nil && tmp1 != nil {
		if toggler%2 == 0 {
			tmp0.Next = tmp2
			// move forward
			tmp2 = tmp2.Next
		} else {
			tmp0.Next = tmp1
			// move forward
			tmp1 = tmp1.Next
		}
		// move tmp2
		tmp0 = tmp0.Next
		// increase toggler
		toggler += 1
	}
	// corner cases
	// if tmp1 still remains, add it to tmp0
	if tmp1 != nil {
		tmp0.Next = tmp1
	}
	// if tmp2 still remains, add it to tmp0
	if tmp2 != nil {
		tmp0.Next = tmp2
	}
	return l1
}

func ZipLinkedListsRecur(l1, l2 *singly.Node[int]) *singly.Node[int] {
	// base cases
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// keep l1.nxt
	nxt0 := l1.Next
	// keep l1.nxt
	nxt1 := l2.Next
	// set l1.nxt = l2
	l1.Next = l2
	l2.Next = ZipLinkedListsRecur(nxt0, nxt1)
	return l1
}
