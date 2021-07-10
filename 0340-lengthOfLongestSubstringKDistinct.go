package leetcode

import "math"

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 || len(s) == 0 {
		return 0
	}
	left, right := 0, 0
	hashMap := make(map[uint8]int)
	var res int
	for right < len(s) {
		hashMap[s[right]] = right
		right++
		if len(hashMap) == k+1 {
			index := math.MaxInt64
			for _, value := range hashMap {
				if value < index {
					index = value
				}
			}
			delete(hashMap, s[index])
			left = index + 1
		}
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
