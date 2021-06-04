package leetcode

func maxProduct(nums []int) int {
	curMax, curMin, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := curMax, curMin
		curMax = max(mx*nums[i], max(nums[i], mn*nums[i]))
		curMin = min(mn*nums[i], min(nums[i], mx*nums[i]))
		res = max(curMax, res)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
