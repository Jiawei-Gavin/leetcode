package leetcode

func respace(dictionary []string, sentence string) int {
	dp := make([]int, len(sentence)+1)
	for i := 1; i <= len(sentence); i++ {
		dp[i] = i
	}
	for i := 1; i <= len(sentence); i++ {
		for j := 0; j < len(dictionary); j++ {
			s := dictionary[j]
			if i-len(s) >= 0 && s == sentence[(i-len(s)):i] {
				if dp[i-len(s)] < dp[i] {
					dp[i] = dp[i-len(s)]
				}
			} else if dp[i] > dp[i-1]+1 {
				dp[i] = dp[i-1] + 1
			}
		}
	}
	return dp[len(sentence)]
}
