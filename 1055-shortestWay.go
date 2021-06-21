package leetcode

func shortestWay(source string, target string) int {
	dp := make([][26]int, len(source)+1)
	for i := 0; i < 26; i++ {
		dp[len(source)][i] = len(source)
	}
	for i := len(source) - 1; i >= 0; i-- {
		c := source[i]
		for j := 0; j < 26; j++ {
			if c == (byte)(j+'a') {
				dp[i][j] = i
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	res := 1
	tmp := 0
	for i := 0; i < len(target); i++ {
		c := target[i]
		if dp[0][c-'a'] == len(source) {
			return -1
		}
		p := dp[tmp][c-'a']
		if p == len(source) {
			res++
			p = dp[0][c-'a']
		}
		tmp = p + 1
	}
	return res
}
