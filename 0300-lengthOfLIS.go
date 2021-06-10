package leetcode

func lengthOfLIS(nums []int) int {
	dp, result := make([]int, len(nums)+1), 0
	dp[0] = 0
	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] = dp[i] + 1
		result = max(result, dp[i])
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
