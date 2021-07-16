package leetcode

func longestOnes(nums []int, k int) int {
	res := 0
	left, right := 0, 0
	count := 0
	for right < len(nums) {
		if nums[right] == 0 {
			count++
		}
		for count > k {
			if nums[left] == 0 {
				count--
			}
			left++
		}
		res = max(res, right-left+1)
		right++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
