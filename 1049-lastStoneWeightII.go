package leetcode

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	dp := make([]bool, sum/2+1)
	dp[0] = true
	for _, weight := range stones {
		for i := sum / 2; i >= weight; i-- {
			dp[i] = dp[i] || dp[i-weight]
		}
	}
	for i := sum / 2; ; i-- {
		if dp[i] {
			return sum - 2*i
		}
	}
}
