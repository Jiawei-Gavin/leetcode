package leetcode

func maxSizeSlices(slices []int) int {
	return max(cal(slices[1:]), cal(slices[:len(slices)-1]))
}

func cal(slices []int) int {
	n := len(slices)
	choose := (n + 1) / 3
	dp := make([][]int, len(slices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, choose+1)
	}
	for i := 1; i <= len(slices); i++ {
		for j := 1; j <= choose && j <= i; j++ {
			if i < 2 {
				dp[i][j] = max(dp[i-1][j], slices[i-1])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-2][j-1]+slices[i-1])
			}
		}
	}
	return dp[len(slices)][choose]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
