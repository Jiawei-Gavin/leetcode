package leetcode

func findTargetSumWays(nums []int, target int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	if target > total || (target+total)%2 == 1 {
		return 0
	}
	tmp := (target + total) / 2
	dp := make([]int, tmp+1)
	dp[0] = 1
	for _, n := range nums {
		for i := tmp; i >= n; i-- {
			dp[i] += dp[i-n]
		}
	}
	return dp[tmp]
}
