package leetcode

func getKthMagicNumber(k int) int {
	a, b, c := 0, 0, 0
	dp := make([]int, k)
	dp[0] = 1
	for i := 1; i < k; i++ {
		dp[i] = min(dp[a]*3, min(dp[b]*5, dp[c]*7))
		if dp[i] == dp[a]*3 {
			a++
		}
		if dp[i] == dp[b]*5 {
			b++
		}
		if dp[i] == dp[c]*7 {
			c++
		}
	}
	return dp[k-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
