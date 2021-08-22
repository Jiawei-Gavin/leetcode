package leetcode

func getMaximumGenerated(n int) int {
	var res int
	if n == 0 {
		return res
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i/2] + i%2*dp[i/2+1]
	}
	for _, v := range dp {
		res = max(res, v)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
