package leetcode

func waysToStep(n int) int {
	dp := []int{1, 2, 4}
	if n <= 3 {
		return dp[n-1]
	}
	for i := 3; i < n; i++ {
		tmp := (dp[0] + dp[1] + dp[2]) % 1000000007
		dp[0], dp[1], dp[2] = dp[1], dp[2], tmp
	}

	return dp[2]
}
