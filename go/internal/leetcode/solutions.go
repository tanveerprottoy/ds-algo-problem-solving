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
