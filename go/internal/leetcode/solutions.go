package leetcode

import (
	"fmt"
	"math"
	"sort"
	"strconv"

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
	You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.
*/
func MergeTwoLinkedLists(list1 *MergerListNode, list2 *MergerListNode) *MergerListNode {
	res := new(MergerListNode)
	// cover corner cases
	if list1 == nil {
		res = list2
	} else if list2 == nil {
		res = list1
	} else {
		// iterate the first linked list
		var tmp0 *MergerListNode
		var tmp1 *MergerListNode
		tmp0 = list1
		tmp1 = list2
		for tmp0 != nil {
			res.Val = tmp0.Val
			// iterate the second
			if tmp1 != nil {
				tmp1 = list2
				res.Nxt = tmp1
				tmp1 = tmp1.Nxt
			} else {
				res.Nxt = tmp0.Nxt
			}
			tmp0 = tmp0.Nxt
		}
	}
	return res
}

func LongestCommonSubsequence(text1 string, text2 string) int {
	return -1
}
