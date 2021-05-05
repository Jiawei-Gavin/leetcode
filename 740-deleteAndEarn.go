package leetcode

// solution1
func deleteAndEarn(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxVal := 0
	for _, val := range nums {
		maxVal = max(maxVal, val)
	}
	all := make([]int, maxVal+1)
	for _, value := range nums {
		all[value]++
	}

	dp := make([]int, maxVal+1)
	dp[1] = all[1] * 1
	dp[2] = max(dp[1], all[2]*2)
	for i := 2; i <= maxVal; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+i*all[i])
	}
	return dp[maxVal]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// solution2
func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, value := range nums {
		maxVal = max(maxVal, value)
	}
	sum := make([]int, maxVal+1)
	for _, value := range nums {
		sum[value] += value
	}
	return rob(sum)
}

func rob(nums []int) int {
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[1], nums[0])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[len(nums)-1]
}
