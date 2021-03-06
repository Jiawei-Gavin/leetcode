package leetcode

func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = max(max(2*(i-2), 2*dp[i-2]), max(3*(i-3), 3*dp[i-3]))
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
