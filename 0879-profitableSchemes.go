package leetcode

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, minProfit+1)
		dp[i][0] = 1
	}
	for i, members := range group {
		earn := profit[i]
		for j := n; j >= members; j-- {
			for k := minProfit; k >= 0; k-- {
				dp[j][k] = (dp[j][k] + dp[j-members][max(0, k-earn)]) % (1e9 + 7)
			}
		}
	}
	return dp[n][minProfit]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
