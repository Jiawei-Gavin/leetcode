package leetcode

// solution1
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[1], nums[0])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[len(nums)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// solution2
func rob(nums []int) int {
	use, notUse := 0, 0
	for i := 0; i < len(nums); i++ {
		use, notUse = notUse+nums[i], max(use, notUse)
	}
	return max(use, notUse)
}
