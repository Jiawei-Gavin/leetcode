package leetcode

// O(n)
func numberOfArithmeticSlices(nums []int) int {
	dp := make([]int, len(nums))
	var res int
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = 1 + dp[i-1]
			res += dp[i]
		}
	}
	return res
}

// O(1)
func numberOfArithmeticSlices(nums []int) int {
	dp := 0
	var res int
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp = 1 + dp
			res += dp
		} else {
			dp = 0
		}
	}
	return res
}
