package leetcode

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	helper := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if s[i-1] == p[j-1] || p[j-1] == '.' {
			return true
		}
		return false
	}
	for i := 0; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if p[j-1] == '*' {
				if helper(i, j-1) {
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			} else if helper(i, j) {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}
