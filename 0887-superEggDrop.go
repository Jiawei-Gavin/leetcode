package leetcode

func superEggDrop(k int, n int) int {
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= k; i++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1
			if dp[i][j] >= n {
				return j
			}
		}
	}
	return n
}
