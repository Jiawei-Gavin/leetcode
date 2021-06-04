package leetcode

import "strings"

func wordBreak(s string, wordDict []string) []string {
	m := make(map[string]bool, len(s)+1)
	for _, word := range wordDict {
		m[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	var result []string
	if !dp[len(s)] {
		return result
	}
	dfs(s, m, []string{}, &result)
	return result
}

func dfs(s string, m map[string]bool, res []string, result *[]string) {
	if len(s) == 0 {
		*result = append(*result, strings.Join(res, " "))
		return
	}
	for i := 1; i <= len(s); i++ {
		if m[s[:i]] {
			dfs(s[i:], m, append(res, s[:i]), result)
		}
	}
}
