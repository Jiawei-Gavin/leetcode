package leetcode

func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[left]&1 == 1 {
			left++
		}
		for left < right && nums[right]&1 == 0 {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}
