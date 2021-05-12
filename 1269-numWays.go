package leetcode

func numWays(steps int, arrLen int) int {
	const mod = 1e9 + 7
	dp := make([][]int, steps+1)
	max := min(steps, arrLen)
	for i := range dp {
		dp[i] = make([]int, max+1)
	}
	dp[0][0] = 1
	for i := 1; i <= steps; i++ {
		for n := 0; n < max; n++ {
			dp[i][n] = dp[i-1][n]
			if n > 0 {
				dp[i][n] = (dp[i][n] + dp[i-1][n-1]) % mod
			}
			if n < max {
				dp[i][n] = (dp[i][n] + dp[i-1][n+1]) % mod
			}
		}
	}
	return dp[steps][0]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
