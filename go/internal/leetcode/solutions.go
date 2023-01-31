package leetcode

import (
	"fmt"
	"math"
	"sort"
)

func LengthOfLongestSubstring(s string) int {
	arrLength := len(s)
	fast, slow := 0, 0
	mapper := make(map[string]string)
	result := 0
	for fast < arrLength {
		value := string(s[fast])
		existingIndex := mapper[value]
		if existingIndex != "" {
			// value exist, break the substring count
			// extend the slow range, this reset the start value
			slow += 1
		} else {
			// value does not exist, continue substring count
			// extend the fast range, this extend the result to next index
			// add the value to set
			// calculate result
			mapper[value] = fmt.Sprintf("%d", fast)
			result = int(math.Max(float64(result), float64((fast-slow)+1)))
			fast += 1
		}
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
