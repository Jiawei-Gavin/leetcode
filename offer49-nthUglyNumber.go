package leetcode

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	a, b, c := 0, 0, 0
	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(min(n2, n3), n5)
		if dp[a]*2 == dp[i] {
			a++
		}
		if dp[b]*3 == dp[i] {
			b++
		}
		if dp[c]*5 == dp[i] {
			c++
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
