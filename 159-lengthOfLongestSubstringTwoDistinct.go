package leetcode

import "math"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) < 3 {
		return len(s)
	}
	hashTable := make(map[byte]int)
	result := 2
	left, right := 0, 0
	for right < len(s) {
		if len(hashTable) < 3 {
			hashTable[s[right]] = right
			right++
		}
		if len(hashTable) == 3 {
			idx := getMapMin(hashTable)
			left = idx + 1
			delete(hashTable, s[idx])
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

func getMapMin(hashTable map[byte]int) int {
	min := math.MaxInt64
	for _, value := range hashTable {
		if value < min {
			min = value
		}
	}
	return min
}
