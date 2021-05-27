package leetcode

func lengthOfLongestSubstring(s string) int {
	dict := map[byte]int{}
	dp := make([]int, len(s))
	result := 0
	for i := range s {
		if i == 0 {
			dp[i], dict[s[i]] = 1, 0
		} else if j, ok := dict[s[i]]; !ok {
			dp[i] = dp[i-1] + 1
		} else if dp[i-1] < i-j {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = i - j
		}
		dict[s[i]] = i
		result = max(dp[i], result)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
