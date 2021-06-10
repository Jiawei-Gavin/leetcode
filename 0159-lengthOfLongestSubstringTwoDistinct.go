package leetcode

import "math"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) < 3 {
		return len(s)
	}
	hashMap := make(map[byte]int)
	result := 2
	left, right := 0, 0
	for right < len(s) {
		if len(hashMap) < 3 {
			hashMap[s[right]] = right
			right++
		}
		if len(hashMap) == 3 {
			idx := getMapMin(hashMap)
			left = idx + 1
			delete(hashMap, s[idx])
		}
		result = max(result, right-left)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMapMin(hashMap map[byte]int) int {
	min := math.MaxInt64
	for _, value := range hashMap {
		if value < min {
			min = value
		}
	}
	return min
}
