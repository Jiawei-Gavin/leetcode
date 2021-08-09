package leetcode

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	if numRows == 0 {
		return dp
	}
	for i := range dp {
		dp[i] = make([]int, i+1)
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			}
		}
	}
	return dp
}
