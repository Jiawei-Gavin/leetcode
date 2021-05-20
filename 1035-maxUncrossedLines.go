package leetcode

func maxUncrossedLines(A []int, B []int) (max int) {
	dp := make([][]int, len(A)+1, len(A)+1)
	for index := range dp {
		dp[index] = make([]int, len(B)+1, len(B)+1)
	}
	for index, a := range A {
		for i, b := range B {
			if a == b {
				dp[index+1][i+1] = dp[index][i] + 1
			} else if dp[index+1][i+1] = dp[index][i+1]; dp[index+1][i+1] < dp[index+1][i] {
				dp[index+1][i+1] = dp[index+1][i]
			}
		}
	}
	return dp[len(A)][len(B)]
}
