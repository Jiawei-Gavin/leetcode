package leetcode

func canBeIncreasing(nums []int) bool {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			count++
			if i > 0 && i < len(nums)-2 {
				if nums[i+1] <= nums[i-1] && nums[i] >= nums[i+2] {
					return false
				}
			}
		}
		if count > 1 {
			return false
		}
	}
	return true
}
