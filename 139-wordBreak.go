package leetcode

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, v := range wordDict {
			if dp[i] {
				break
			}
			if i >= len(v) {
				if v == s[i-len(v):i] {
					dp[i] = dp[i-len(v)]
				}
			}
		}
	}
	return dp[len(s)]
}
