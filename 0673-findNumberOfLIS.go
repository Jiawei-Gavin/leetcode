package leetcode

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	count := make([]int, len(nums))
	dp[0] = 1
	count[0] = 1
	maxLength := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}
		maxLength = max(maxLength, dp[i])
	}
	var res int
	for i := 0; i <= len(dp)-1; i += 1 {
		if dp[i] == maxLength {
			res += count[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
