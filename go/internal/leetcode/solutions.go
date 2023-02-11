package leetcode

import (
	"fmt"
	"math"
	"sort"

	"github.com/tanveerprottoy/ds-algo-problem-solving/pkg/adapter"
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
