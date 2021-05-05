package leetcode

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob1(nums[1:]), rob1(nums[:len(nums)-1]))
}

func rob1(nums []int) int {
	use, notUse := 0, 0
	for _, v := range nums {
		use, notUse = notUse+v, max(use, notUse)
	}
	return max(use, notUse)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
