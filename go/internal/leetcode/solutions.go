package leetcode

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/linkedlist/singly"
	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/linkedlist/singly/generic"
	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/stack"
	"github.com/tanveerprottoy/ds-algo-problem-solving/pkg/adapter"
	"github.com/tanveerprottoy/ds-algo-problem-solving/pkg/slice"
)

func LengthOfLongestSubstring(s string) int {
	slow := 0
	mapper := make(map[string]string)
	result := 0
	for fast := range s {
		value := string(s[fast])
		existingIndex := mapper[value]
		if existingIndex != "" {
			// value exists, break the substring count
			// extend the slow range, this reset the start value
			existingIndexNum, err := adapter.StringToFloat(existingIndex, 64)
			if err != nil {
				continue
			}
			slow = int(math.Max(existingIndexNum, float64(slow)))
		}
		// value does not exist, continue substring count
		// calculate result
		// add the value to map
		result = int(math.Max(float64(result), float64((fast-slow)+1)))
		// store in map with index starting from 1
		// as it ensures the calcaulation (fast-slow)+1 will be correct
		mapper[value] = fmt.Sprintf("%d", fast+1)
	}
	return result
}

// ineffiecient solution
func MedianOfTwoSortedArray(nums1, nums2 []int) float64 {
	if nums1 == nil || nums2 == nil {
		return 0.0
	}
	mergedArr := append(nums1, nums2...)
	sort.Ints(mergedArr)
	mergedArrLength := len(mergedArr)
	midPoint := int((mergedArrLength + 1) / 2)
	median := 0.0
	if mergedArrLength%2 == 0 {
		// for even length, median is avg of midPoint - 1 + midPoint elements
		// median = (mergedArr[midPoint - 1] + mergedArr[midPoint]) / 2
		median = float64((mergedArr[midPoint-1] + mergedArr[midPoint])) / 2.0
	} else {
		median = float64(mergedArr[midPoint-1])
	}
	return median
}

func LongestPalindrome(s string) string {
	l := len(s)
	if l == 0 || l == 1 {
		return s
	}
	res := ""
	tmp := ""
	start, end := 0, l-1
	for start < l && end >= 0 {
		val := s[start]
		if val == s[end] {
			tmp += string(val)
		} else {
			// palindrome pattern broke, reset
			if len(tmp) > len(res) {
				res = tmp
			}
			tmp = string(val)
		}
	}
	return res
}

/*
	Given an array of integers arr and an integer k.

Find the least number of unique integers after removing exactly k elements.
Example 1:
Input: arr = [5,5,4], k = 1
Output: 1
Explanation: Remove the single 4, only 5 is left.
Example 2:
Input: arr = [4,3,1,1,3,3,2], k = 3
Output: 2
Explanation: Remove 4, 2 and either one of the two 1s or three 3s. 1 and 3 will be left.

Constraints:

1 <= arr.length <= 10^5
1 <= arr[i] <= 10^9
0 <= k <= arr.length
*/
func LeastNumOfUniqueInts(arr []int, k int) int {
	data := make(map[int]int)
	for i := range arr {
		elem := arr[i]
		count := data[elem]
		if count != 0 {
			// value exists, increase count
			count++
			data[elem] = count
			continue
		}
		data[elem] = 1
	}
	// get the values
	var values []int
	for _, v := range data {
		values = append(values, v)
	}
	// sort the counts
	sort.Ints(values)
	// remove k elements
	// run the loop till i < len(values) or k <= 0
	// k is decreased inside
	for i := 0; i < len(values); i++ {
		if k <= 0 {
			break
		}
		c := values[i]
		if k >= c {
			values = slice.RemoveAt(values, i)
			// as item del from slice, decrease i
			// so that the next iteration happens
			// otherwise
			i--
		} else {
			c -= k
			values[i] = c
		}
		k -= c
	}
	return len(values)
}

func LeastNumOfUniqueInts1(arr []int, k int) int {
	data := make(map[int]int)
	for i := range arr {
		elem := arr[i]
		count := data[elem]
		if count != 0 {
			// value exists, increase count
			count++
			data[elem] = count
			continue
		}
		data[elem] = 1
	}
	// get the values
	var values []int
	for _, v := range data {
		values = append(values, v)
	}
	// sort the counts
	sort.Ints(values)
	// remove k elements
	// run the loop till i < len(values) or k <= 0
	// k is decreased inside
	for i := 0; i < len(values); i++ {
		c := values[i]
		// iterate over the map to find item for value
		for key, val := range data {
			if val == c {
				// item found
				if k >= c {
					// ex: k = 3, c = 3, k >= c == true
					// remove item as it's count is <= k
					delete(data, key)
				} else {
					// ex: k = 2, c = 4, k >= c == false,
					// c -= k == 2
					// as c > k, c = c - k,
					// then store it in the map
					c -= k
					data[key] = c
				}
				// subtract from k = k -c
				// so that k is updated according to count
				k -= c
				break
			}
		}
		if k <= 0 {
			break
		}
	}
	return len(data)
}

/*
	Given a signed 32-bit integer x, return x with its digits reversed.

If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
Example 1:

Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21

//pop operation:
pop = x % 10;
x /= 10;

//push operation:
temp = rev * 10 + pop;
rev = temp;
*/
func ReverseInteger(x int) int {
	rev := 0
	var mod int
	for x != 0 {
		// pop
		mod = x % 10
		x /= 10
		/*	prevent overflow
			for positive num
			if rev > (math.MaxInt32/10) then rev := rev*10 + mod will overflow
			if rev == (math.MaxInt32/10) then will overflow if and only if mod > 7
		*/
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && mod > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && mod < -8) {
			return 0
		}
		// push
		rev = rev*10 + mod
	}
	return rev
}

/*
	Given an array of integers arr, replace each element with its rank.

The rank represents how large the element is. The rank has the following rules:

Rank is an integer starting from 1.
The larger the element, the larger the rank. If two elements are equal, their rank must be the same.
Rank should be as small as possible.

Example 1:

Input: arr = [40,10,20,30]
Output: [4,1,2,3]
Explanation: 40 is the largest element. 10 is the smallest. 20 is the second smallest. 30 is the third smallest.
*/
func ArrayRankTransform(arr []int) []int {
	data := make(map[int]int)
	var arrCopy []int
	var res []int
	rank := 0
	arrCopy = slice.AppendCopy(arr, arrCopy)
	sort.Ints(arr)
	for i := range arr {
		val := arr[i]
		if i > 0 {
			if val == arr[i-1] {
				// keep the last rank
				// no need to insert in map
				continue
			}
		}
		rank++
		data[val] = rank
	}
	for i := range arrCopy {
		val := arrCopy[i]
		rank := data[val]
		res = append(res, rank)
	}
	return res
}

func IsPalindromeNumber(x int) bool {
	str := strconv.Itoa(x)
	len := len(str)
	if len == 1 {
		return true
	}
	for i, j := 0, len-1; i < len/2; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

/*
	Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.

Example 1:

Input: s = "III"
Output: 3
Explanation: III = 3.
*/
func RomanToInt(s string) int {
	romans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	len := len(s)
	res := 0
	// iterate from 0 to the item before last
	// in this approach the last element has to be added to result
	// as the last one will not be reached for nxt < len(s)
	for curr, nxt := 0, 1; nxt < len; curr, nxt = curr+1, nxt+1 {
		currVal := romans[string(s[curr])]
		nxtVal := romans[string(s[nxt])]
		// as roman numerals are written from largest to smallest from left to right
		// left(currVal) should be larger or equal to right(nxtVal) to be added up in res
		// otherwise subtract currVal from res
		// if current is larger than or equal nxt
		// add to res
		// else subtract it from rest
		if currVal >= nxtVal {
			res += currVal
		} else {
			res -= currVal
		}
	}
	// need to add the last item to res
	// as the last one will did not be reach for nxt < len(s)
	return res + romans[string(s[len-1])]
}

/*
	Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:
*/
// vertical scan solution
func LongestCommonPrefixVertical(strs []string) string {
	res := ""
	arrLength := len(strs)
	if arrLength == 0 {
		return res
	}
	// iterate for the first element of strs
	// stop and return result when strs[0]
	// has been iterated
	for i := range strs[0] {
		char := string(strs[0][i])
		for j := 1; j < arrLength; j++ {
			// iterate over other elements after strs[0]
			// it will only iterate len(strs) - 1 time as j < l
			// here j goes over other elements other than strs[0]
			// and i used to iterate over char of string
			// need to stop if i == other elments length
			// stop if i == len(strs[j])
			// or char does not match with target
			if i == len(strs[j]) || char != string(strs[j][i]) {
				return res
			}
		}
		res += char
	}
	return res
}

// horizontal scan solution
func LongestCommonPrefixHorizontal(strs []string) string {
	res := ""
	arrLen := len(strs)
	if arrLen == 0 {
		return res
	}
	res = strs[0]
	// iterate over the other elements strs[n],
	// where n > 0
	for i := 1; i < arrLen; i++ {
		// get the [i]th str
		str := strs[i]
		// get the length that can be used for str
		// as len(str) might be smaller or larger
		// than strs[0]
		adjustedLen := len(res)
		if adjustedLen > len(str) {
			adjustedLen = len(str)
		}
		// in this we are removing char from
		// strs[0] until it mathes with some or
		// "" incase no match found
		// need to loop until str[:len(res)] != res
		// get only the string part that is
		// equal in length of strs[0]
		// which is str[:len(res)]
		// incase of len(strs[0]) > len(str)
		// take the length from len(str)
		// other breaking condition will be inside the loop
		for str[:adjustedLen] != res {
			// remove char from res = strs[0]
			// remove 1 char from last
			res = res[:len(res)-1]
			// check if str = ""
			// then return res
			if res == "" {
				return res
			}
			// get len(res) after removing
			// if len(res) <= adjustedLen
			// reassign adjustedLen = len(res)
			// as we want to match the word with strs[0]
			// this ensures index out of range exception
			// not happens and str[:end] also calculated correctly
			// want to start the loop with len(str)
			// in case len(res) > len(str)
			// then after removing item[s]
			// when len(res) <= len(str)
			// will set adjustedLen = len(res)
			resLen := len(res)
			if resLen <= adjustedLen {
				adjustedLen = resLen
			}
		}
	}
	return res
}

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/
func IsValidParentheses(s string) bool {
	// put brackets
	bracketsMap := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	// init the stack
	brackets := stack.NewStack[string]()
	for i := range s {
		char := string(s[i])
		// check if its opening or
		// close bracket
		brac := bracketsMap[char]
		if brac != "" {
			// it's opening bracket
			brackets.Push(char)
		} else if brackets.IsEmpty() {
			// closed bracket was found
			// but the stack is empty
			// no open bracket was added
			// in this case its invalid
			return false
		} else {
			// pop from stack
			p := brackets.PopAlt()
			// popped is the
			closeBrac := bracketsMap[p]
			fmt.Println(p)
			fmt.Println(closeBrac)
			if char != closeBrac {
				return false
			}
		}
	}
	return brackets.IsEmpty()
}

/*
Given an integer array nums and an integer val, remove all occurrences of val
in nums in-place. The relative order of the elements may be changed.
Since it is impossible to change the length of the array in some languages,
you must instead have the result be placed in the first part of the array nums.
More formally, if there are k elements after removing the duplicates,
then the first k elements of nums should hold the final result.
It does not matter what you leave beyond the first k elements.
Return k after placing the final result in the first k slots of nums.
Do not allocate extra space for another array. You must do this by modifying
the input array in-place with O(1) extra memory.

Example 1:

Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of
nums being 2.
It does not matter what you leave beyond the returned k
(hence they are underscores).
*/
func RemoveElement(nums []int, val int) int {
	// effiecient 2 pointers solution
	// 2 pointer technique is keeping
	// a 2nd pointer along with the
	// looping one which tracks the
	// index that will be overwritten
	if len(nums) == 0 {
		return 0
	}
	// 2nd index
	// k tracks the last inserted
	// non target val nums[i] != val
	k := 0
	for i := 0; i < len(nums); i++ {
		// move values that are not
		// the target(val)
		// from start of array
		// keep increasing k whenever
		// nums[i] != val
		if nums[i] != val {
			// when n != val, want to
			// move the item to k
			// modIdx starts from 0 then
			// increments only when element
			// is not equal to val
			nums[k] = nums[i]
			k++
		}
	}
	fmt.Println("nums", nums)
	return k
}

/*
Given an integer array nums sorted in non-decreasing order,
remove the duplicates in-place such that each unique element appears only once.
The relative order of the elements should be kept the same.
Since it is impossible to change the length of the array in some languages,
you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
Return k after placing the final result in the first k slots of nums.
Do not allocate extra space for another array. You must do this by modifying
the input array in-place with O(1) extra memory.

Example 1:
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements
of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k
(hence they are underscores).
*/
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 2nd index
	// k tracks the last inserted
	// unique val
	// start from 1 as
	// comparision needed from i = 1
	k := 1
	// similar as RemoveElement
	// here duplicates will be removed
	// start from index 1
	for i := 1; i < len(nums); i++ {
		// as nums arr is sorted
		// duplicates will be one after another
		// need to check with the previous item
		if nums[i] != nums[i-1] {
			// if not equal set nums[k] = nums[i]
			// increase k after move
			nums[k] = nums[i]
			k++
		}
	}
	fmt.Println("nums", nums)
	return k
}

func RemoveDuplicatesMap(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	uniques := make(map[int]int)
	k := 0
	// similar as RemoveElement
	// here duplicates will be removed
	for i := 0; i < len(nums); i++ {
		// check if value exists in map
		if _, ok := uniques[nums[i]]; ok {
			// it's not in map, move to k
			// increase k after move
			nums[k] = nums[i]
			k++
		}
		// add val to map
		uniques[nums[i]] = i
	}
	fmt.Println("nums", nums)
	fmt.Println("uniques", uniques)
	return k
}

/*
Given an integer array nums sorted in non-decreasing order,
remove some duplicates in-place such that each unique element
appears at most twice. The relative order of the elements should be kept the same.
Since it is impossible to change the length of the array in some languages,
you must instead have the result be placed in the first part of the array nums.
More formally, if there are k elements after removing the duplicates, then the
first k elements of nums should hold the final result. It does not matter what
you leave beyond the first k elements.
Return k after placing the final result in the first k slots of nums.
Do not allocate extra space for another array. You must do this by modifying
the input array in-place with O(1) extra memory.

Example 1:
Input: nums = [1,1,1,2,2,3]
Output: 5, nums = [1,1,2,2,3,_]
Explanation: Your function should return k = 5, with the first five elements
of nums being 1, 1, 2, 2 and 3 respectively.
It does not matter what you leave beyond the returned k
(hence they are underscores).
*/
func RemoveDuplicatesOverTwo(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// frequency map
	freq := make(map[string]string)
	k := 0
	// similar as RemoveDuplicates
	// here duplicate count of over 2 will be removed
	for i := 0; i < len(nums); i++ {
		// check if value exists in map
		if freq[fmt.Sprint(nums[i])] == "" {
			// if not in map, move to k
			// increase k after move
			nums[k] = nums[i]
			k++
			// add val to map with frequency
			freq[fmt.Sprint(nums[i])] = fmt.Sprint(1)
		} else {
			// val exists in map
			// check the frequency of it
			// if frequncy is < 2 move to k
			// if frequncy is > 2 do nothing
			f, _ := strconv.Atoi(freq[fmt.Sprint(nums[i])])
			if f < 2 {
				nums[k] = nums[i]
				k++
				f++
				freq[fmt.Sprint(nums[i])] = fmt.Sprint(f)
			}
		}
	}
	fmt.Println("nums", nums)
	return k
}

/*
	Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.
Example 1:
Input: nums = [1,3,5,6], target = 5
Output: 2
*/
func SearchInsert(nums []int, target int) int {
	i := 0
	for i < len(nums) {
		// if match found return i
		if nums[i] == target {
			return i
		} else if i > 0 {
			// if iterated to atleast 2nd item
			// check with current and previous
			// target needs to be => prev and <= curr
			// in order to sit in i
			if target >= nums[i-1] && target <= nums[i] {
				return i
			}
		} else {
			// in case when i == 0
			// need to check if target <= nums[i]
			// in this case return 0
			if target < nums[i] {
				return i
			}
		}
		// increment i as loop iterates to next
		i++
	}
	return i
}

/*
	Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal
substring

	consisting of non-space characters only.

Example 1:
Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.
*/
func LengthOfLastWord(s string) int {
	// remove right/end whitespaces
	s = strings.TrimRight(s, " ")
	i := len(s) - 1
	// iterate from last
	for i >= 0 {
		// look for space
		if fmt.Sprintf("%c", s[i]) == " " {
			// space found word ends
			// as iterating from last
			// as i is set as len(s) - 1
			// and expected to loop till 0
			// and stop when -1
			// need to add the 1 in order to get
			// the correct length
			// return len(s) - (i + 1)
			return len(s) - (i + 1)
		}
		// decrement and continue
		i--
	}
	// has to be (i + 1) as the loop exists
	// when i == -1
	return len(s) - (i + 1)
}

/*
	You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.

Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
*/
func PlusOne(digits []int) []int {
	// iterate from last
	for i := len(digits) - 1; i >= 0; i-- {
		// check if digits[i] < 9
		// in this case task is simple
		// just add 1 to digits[i]
		if digits[i] < 9 {
			// can add
			// and return the array
			// this case handles the last
			// element and also the carry 1
			// case which occurs digits[i] == 9
			digits[i]++
			return digits
		}
		// digits[i] == 9
		// make curr elem digits[i] = 0
		// this creates a carry 1 value
		digits[i] = 0
	}
	// if the code reached here
	// it means has a carry 1
	// the array has to be increased
	// in size by 1 and 1 to be added in
	// newArr[0] array
	var newArr []int
	newArr = append(newArr, 1)
	newArr = append(newArr, digits...)
	return newArr
}

func IsPalindromeAfterRemoval(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	s = strings.ToLower(s)
	r := regexp.MustCompile(`[^a-zA-Z0-9]`)
	s = r.ReplaceAllString(s, "")
	fmt.Println(s)
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		if fmt.Sprintf("%c", s[i]) != fmt.Sprintf("%c", s[j]) {
			return false
		}
	}
	return true
}

/*
	Given an integer array nums, return true if any value appears at least twice in the array,
	and return false if every element is distinct.

Example 1:

Input: nums = [1,2,3,1]
Output: true
*/
func ContainsDuplicate(nums []int) bool {
	uniques := make(map[int]int)
	for i := range nums {
		// check if value is in map
		if _, ok := uniques[nums[i]]; ok {
			// if its in map
			// its duplicate return true
			return true
		}
		// add val to map
		uniques[nums[i]] = i
	}
	// no duplicates found
	return false
}

/*
	Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.
Example 1:
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
*/
func MoveZeroes(nums []int) []int {
	// two index solution
	// 2nd index
	k := 0
	for i := range nums {
		// check if not zero
		if nums[i] != 0 {
			// move non zero to k,
			// swap values, set val of k = i
			// and val of prev i = prev k
			// assign nums[k] = nums[i]
			// and prev nums[i] = prev nums[k]
			// store current nums[k]
			tmp := nums[k]
			// swap nums[k] = nums[i]
			// then set nums[i] = tmp, which is previous nums[k]
			// this way nums[i] will retain the overwritten velues
			// as this algo only swaps the non zero values
			// the retained value will be 0
			nums[i] = tmp
			// done in go's shorthand way
			// nums[k], nums[i] = nums[i], nums[k]
			// increment k
			k++
		}
		// no need to anything when nums[i] = 0
	}
	return nums
}

/*
	Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times.
You may assume that the majority element always exists in the array.
*/
func MajorityElement(nums []int) int {
	// check corener cases
	if len(nums) == 0 {
		return 0
	}
	// if nums has only 1 item
	// return that
	if len(nums) == 1 {
		return nums[0]
	}
	// frequency map
	freq := make(map[int]int)
	for i := range nums {
		// check if item is in map
		if val, ok := freq[nums[i]]; ok {
			// increase freq
			val++
			// check if freq > len(nums)/2
			if val > len(nums)/2 {
				return nums[i]
			}
			// else store the updated freq
			freq[nums[i]] = val
		} else {
			freq[nums[i]] = 1
		}
	}
	return 0
}

/*
	Given a non-empty array of integers nums,

every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

Example 1:

Input: nums = [2,2,1]
Output: 1
*/
func SingleNumber(nums []int) int {
	// store items in a map
	freqs := make(map[int]int)
	for i := range nums {
		if freq, ok := freqs[nums[i]]; ok {
			// if item exists in map
			// check freq == 1
			// as every element appears twice except for one
			if freq == 1 {
				// if freq == 1
				// twice appeared item it is
				// del this it will ensure
				// freq == 1 will remain
				delete(freqs, nums[i])
			} else {
				// increase freq
				freq++
				// store in map
				freqs[nums[i]] = freq
			}
		} else {
			freqs[nums[i]] = 1
		}
	}
	for k := range freqs {
		// return the first key
		return k
	}
	return 0
}

/*
	Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

Example 1:
Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.
*/
func MissingNumber(nums []int) int {
	// first sort the nums
	sort.Ints(nums)
	// loop from 0 to len(nums)
	for i := range nums {
		// in each iteration
		// should be i == nums[i]
		// if not then it's
		// the missing number
		// is i return it
		if i != nums[i] {
			return i
		}
	}
	return -1
}

/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together
the nodes of the first two lists.

Return the head of the merged linked list.
*/
func MergeTwoLists(list1, list2 *singly.Node) *singly.Node {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	// check which has the smaller value at first,
	// that will be the starting list
	// will mutate the starting list, store & traverse with tmp0
	// result will be starting list
	var tmp0 *singly.Node
	var tmp1 *singly.Node
	var tmp2 *singly.Node
	var res *singly.Node
	if list1.Val.(int) <= list2.Val.(int) {
		tmp0 = list1
		// list1 traversal, starting point must be
		// next as tmp0 is assigned to l1
		tmp1 = list1.Next
		// list2 traversal
		tmp2 = list2
		res = list1
	} else {
		tmp0 = list2
		// list2 traversal, starting point must be
		// next as tmp0 is assigned to l2
		tmp1 = list2.Next
		// list1 traversal
		tmp2 = list1
		res = list2
	}
	// loop till one of them is null
	for tmp1 != nil && tmp2 != nil {
		// check and compare both values
		if tmp1.Val.(int) <= tmp2.Val.(int) {
			tmp0.Next = tmp1
			// move to next
			tmp1 = tmp1.Next
		} else if tmp2.Val.(int) <= tmp1.Val.(int) {
			tmp0.Next = tmp2
			// move to next
			tmp2 = tmp2.Next
		}
		// move list1 forward
		tmp0 = tmp0.Next
	}
	// handle if any remaining items are there
	if tmp1 != nil {
		tmp0.Next = tmp1
	}
	if tmp2 != nil {
		tmp0.Next = tmp2
	}
	return res
}

/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
and two integers m and n, representing the number of elements in nums1 and nums2
respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be
stored inside the array nums1. To accommodate this, nums1 has a length of m + n,
where the first m elements denote the elements that should be merged, and the last
n elements are set to 0 and should be ignored. nums2 has a length of n.
*/
func MergeSortedArraysSimple(nums1 []int, m int, nums2 []int, n int) {
	// need to mutate and decorate nums1 as the result
	if m == 0 {
		nums1 = nums2
		return
	}
	if n == 0 {
		return
	}
	// first pointer to remember
	// the last index visited on the 1st arr
	i := 0
	// will keep a second pointer to remember
	// the last index visited on the 2nd arr
	j := 0
	tmp := []int{}
	for i+j < m+n {
		// only operate on nums1 when i < m
		if i < m && nums1[i] <= nums2[j] {
			// keep the nums1[i] element
			tmp = append(tmp, nums1[i])
			// increment i
			i++
		} else {
			// only operate on nums2 when j < n
			if j < n {
				// keep the nums2[j] element
				tmp = append(tmp, nums2[j])
				// increment j
				j++
			}
		}
	}
	nums1 = tmp
	fmt.Println(nums1)
}

/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
and two integers m and n, representing the number of elements in nums1 and nums2
respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be
stored inside the array nums1. To accommodate this, nums1 has a length of m + n,
where the first m elements denote the elements that should be merged, and the last
n elements are set to 0 and should be ignored. nums2 has a length of n.
*/
func MergeSortedArrays(nums1 []int, m int, nums2 []int, n int) {
	// need to mutate and decorate nums1 as the result
	if m == 0 {
		nums1 = nums2
		fmt.Println(nums1)
		return
	}
	if n == 0 {
		return
	}
	// will start from the last of num1
	// as it will be easier to shift from last
	// will use 3 pointers
	// i to keep track of nums1
	// j to keep track of nums2
	// k to keep track of last of nums1 where
	// inserting values will be started
	// all of them will start from last
	i, j, k := m-1, n-1, (m+n)-1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			// store nums1[i] on k
			nums1[k] = nums1[i]
			// move i 1 place backward
			i--
		} else {
			nums1[k] = nums2[j]
			// move j 1 place backward
			j--
		}
		// move k 1 place backward
		k--
	}
	fmt.Println(nums1)
}

func MergeSortedArraysForDLC(nums1 []int, m int, nums2 []int, n int) {
	// need to mutate and decorate nums1 as the result
	// will start from the last of num1
	// as it will be easier to shift from last
	// will use 3 pointers
	// i to keep track of nums1
	// j to keep track of nums2
	// k to keep track of last of nums1 where
	// inserting values will be started
	// all of them will start from last
	i, j, k := m-1, n-1, (m+n)-1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			// store nums1[i] on k
			nums1[k] = nums1[i]
			// move i 1 place backward
			i--
		} else {
			nums1[k] = nums2[j]
			// move j 1 place backward
			j--
		}
		// move k 1 place backward
		k--
	}
	// will need to check for leftovers
	// as for i >= 0 the loop will break
	// j might still be j >= 0
	for j >= 0 {
		// store the remainder of nums2
		nums1[k] = nums2[j]
		j--
		k--
	}
	fmt.Println(nums1)
}

/*
Given an integer array nums sorted in non-decreasing order, return an array of the squares
of each number sorted in non-decreasing order. Example 1:
Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].
Example 2:
Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]
*/
func SortedSquaresSimple(nums []int) []int {
	for i := range nums {
		// calculate and store the sqaure of the element
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}

/*
Given an integer array nums sorted in non-decreasing order, return an array of the squares
of each number sorted in non-decreasing order. Example 1:
Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].
Example 2:
Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]
*/
func SortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	// will use 2 pointers to the left
	// and to the right
	// initial smallest side
	i := 0
	// initial largest side
	j := len(nums) - 1
	// new array tracker
	// which will add element
	// from last
	k := j
	// iterate till left <= right
	for i <= j {
		// create the square first
		// for both ends
		// compare them and set the largest
		// val to the right
		if nums[i]*nums[i] >= nums[j]*nums[j] {
			res[k] = nums[i] * nums[i]
			// increment i
			i++
		} else {
			res[k] = nums[j] * nums[j]
			// decrement j
			j--
		}
		// decrement k res tracker index
		k--
	}
	return res
}

/*
Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
Return the linked list sorted as well.
ex1: Input: head = [1,1,2]
Output: [1,2]
ex2: Input: head = [1,1,2,3,3]
Output: [1,2,3]
*/
func DeleteDuplicates(head *generic.Node[int]) *generic.Node[int] {
	// corner case
	if head == nil || head.Next == nil {
		return head
	}
	// will mutate the head
	// result will be head
	tmp := head
	// traverse the list
	for tmp != nil && tmp.Next != nil {
		// as the list is sorted check with
		// next val
		// if they are equal remove the next
		if tmp.Val == tmp.Next.Val {
			// will remove the tmp.next reference
			// from tmp, and set tmp.next to tmp.next.next
			tmp.Next = tmp.Next.Next
			// need to continue the loop, as the newly changed
			// tmp.next might be equal to tmp val
		} else {
			// will only move to next node when 2 adjacent nodes
			// do not match, as the current node might match with
			// with the changed next node, need to check it first
			// move to next node
			tmp = tmp.Next
		}
	}
	return head
}

/*
Given the head of a sorted linked list, delete all nodes that have duplicate numbers,
leaving only distinct numbers from the original list. Return the linked list sorted as well.
*/
func DeleteDuplicates2(head *generic.Node[int]) *generic.Node[int] {
	// corner case
	if head == nil || head.Next == nil {
		return head
	}
	// store the distinct node
	dummy := *generic.NewNode(0, head)
	// dummy := head
	// distinct value pointer
	tmp := &dummy
	// will return dummy, so will traverse with
	// the head, no need for tmp reference
	// traverse the list
	for head != nil && head.Next != nil {
		// as the list is sorted check with
		// next val
		if head.Val == head.Next.Val {
			// traverse & look for further non distinct values
			for head != nil && head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			// next distinct node found
			// link to dummy
			tmp.Next = head.Next
		} else {
			// move dummy pointer, as it's distinct node
			tmp = tmp.Next
		}
		// move head
		head = head.Next
	}
	return dummy.Next
}

/*
Given two binary strings a and b, return their sum as a binary string.
Example 1:
Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
*/
func AddBinary(a string, b string) string {
	if len(b) == 0 {
		return a
	}
	if len(a) == 0 {
		return b
	}
	carr := 0
	res := ""
	// will use two pointers
	// i = len(a) - 1
	// j = len(b) - 1
	i := len(a) - 1
	j := len(b) - 1
	// loop till both of them are not >= 0
	for i >= 0 || j >= 0 {
		// need to convert the str to int
		// build up from there
		// initial value will be 0
		val0 := 0
		val1 := 0
		if i >= 0 {
			// only access array if it's index is valid
			val0, _ = strconv.Atoi(string(a[i]))
		}
		if j >= 0 {
			// only access array if it's index is valid
			val1, _ = strconv.Atoi(string(b[j]))
		}
		// max possible value
		// 1 + 1 + 1 = 3
		total := val0 + val1 + carr
		// digit will be mod 2
		// as its binary
		// 3 % 2 = 1
		// 2 % 2 = 0
		// 1 % 2 = 1
		digit := total % 2
		// prepend to res
		res = fmt.Sprint(digit) + res
		// carry will be divided by 2
		// as its binary
		// 3 / 2 = 1.5 -> int 1
		// 2 / 2 = 1
		// 1 / 2 = 0.5 -> int 0
		carr = total / 2
		i--
		j--
	}
	// carry might be 1 after the loop
	// check and add if its 1
	if carr > 0 {
		res = "1" + res
	}
	return res
}

/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.
There is a cycle in a linked list if there is some node in the list that can be
reached again by continuously following the next pointer. Internally, pos is used
to denote the index of the node that tail's next pointer is connected to.
Note that pos is not passed as a parameter.
Return true if there is a cycle in the linked list. Otherwise, return false.
*/
func HasCycleSimple(head *generic.Node[int]) bool {
	// the linked list is circular when
	// If any node seems to be pointing towards the head or starting node
	// If no node is pointing to null.
	tmp := head
	// iterate till pointer is not null
	for tmp != nil {
		// check if next node points to head
		if tmp.Next == head {
			// it has completed a cycle, has cycle
			return true
		}
		tmp = tmp.Next
	}
	return false
}

/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.
There is a cycle in a linked list if there is some node in the list that can be
reached again by continuously following the next pointer. Internally, pos is used
to denote the index of the node that tail's next pointer is connected to.
Note that pos is not passed as a parameter.
Return true if there is a cycle in the linked list. Otherwise, return false.
*/
func HasCycle(head *generic.Node[int]) bool {
	// Floyd's Algorithm
	// Also called Tortoise and hare algorithm
	// the linked list is circular when
	// fast & and slow meets at certain step
	// fast will catch up to slow eventually if
	// the linked list has cycle
	fast, slow := head, head
	// iterate till fast != nil && fast.Next != nil && slow != nil
	// it will only stop when
	for fast != nil && fast.Next != nil && slow != nil {
		// slow 1 step
		slow = slow.Next
		// fast 2 steps
		fast = fast.Next.Next
		// check if fast and slow are equal
		if fast == slow {
			// cycle found
			return true
		}
	}
	// cycle was not found
	// loop finished
	return false
}

func LongestCommonSubsequence(text1 string, text2 string) int {
	return -1
}
