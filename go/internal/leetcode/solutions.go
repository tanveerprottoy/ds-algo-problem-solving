package leetcode

import (
	"math"
	"sort"
)

func LengthOfLongestSubstring(str string) {
	arrLength := len(str)
	fast := 0
	slow := 0
	mapper := make(map[string]string)
	result := 0
	for fast < arrLength && slow < arrLength {
		value := string(str[fast])
		existingIndex := mapper[value]
		if existingIndex != "" {
			// value exist, break the substring count
			// extend the slow range, this reset the start value
			// remove the existing value from map
			delete(mapper, mapper[value])
			slow += slow
		} else {
			// value does not exist, continue substring count
			// extend the fast range, this extend the result to next index
			// add the value to set
			// calculate result
			fast += 1
			mapper[value] = string(fast)
			result = int(math.Max(float64(result), float64((fast-slow)+1)))
		}
	}
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
