package leetcode

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			res = min(res, dp[i-j*j])
		}
		dp[i] = res + 1
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
