package leetcode

import "math"

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
