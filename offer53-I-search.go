package leetcode

func search(nums []int, target int) int {
	return helper(nums, target) - helper(nums, target-1)
}

func helper(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
