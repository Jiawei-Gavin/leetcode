package leetcode

func longestValidParentheses(s string) int {
	var res int
	dp := make([]int, len(s)+1)
	for i := 1; i < len(s); i++ {
		if s[i] == ')' && i-dp[i]-1 >= 0 && s[i-dp[i]-1] == '(' {
			dp[i+1] = dp[i] + 2 + dp[i-dp[i]-1]
		}
		res = max(res, dp[i+1])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
