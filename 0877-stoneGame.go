package leetcode

// solution1
func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = piles[i]
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[j] = max(piles[i]-dp[j], piles[j]-dp[j-1])
		}
	}
	return dp[n-1] > 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// solution2 ---- 摆烂之王
func stoneGame(piles []int) bool {
	return true
}
