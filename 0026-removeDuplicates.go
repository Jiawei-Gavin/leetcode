package leetcode

func removeDuplicates(nums []int) int {
	result := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[result] = nums[i+1]
			result++
		}
	}
	return result
}
