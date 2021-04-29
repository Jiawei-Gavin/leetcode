package leetcode

// solution1
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	bitset := make([]bool, 256)
	result, right, left := 0, 0, 0
	for left < len(s) {
		if bitset[s[right]] {
			bitset[s[left]] = false
			left++
		} else {
			bitset[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}

		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

// solution2
func lengthOfLongestSubstring(s string) int {
	hashTable := make(map[byte]int)
	result := 0
	for start, end := 0, 0; end < len(s); end++ {
		if p, ok := hashTable[s[end]]; ok {
			start = max(p, start)
		}
		hashTable[s[end]] = end + 1
		result = max(result, end-start+1)

	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
