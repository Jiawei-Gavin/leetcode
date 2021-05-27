package leetcode

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if j == 0 {
				dp[i][j] = i == 0
			} else {
				if p[j-1] != '*' {
					if i >= 1 && (p[j-1] == '.' || s[i-1] == p[j-1]) {
						dp[i][j] = dp[i-1][j-1]
					}
				} else {
					if j >= 2 {
						dp[i][j] = dp[i][j-2] || dp[i][j]
					}
					if (j >= 2 && i >= 1) && (s[i-1] == p[j-2] || p[j-2] == '.') {
						dp[i][j] = dp[i-1][j] || dp[i][j]
					}
				}
			}
		}

	}
	return dp[m][n]
}
