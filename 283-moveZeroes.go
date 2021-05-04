package leetcode

func moveZeroes(nums []int) {
	j, zero := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zero++
		} else {
			nums[j] = nums[i]
			j++
		}
	}
	for i := len(nums) - zero; i < len(nums); i++ {
		nums[i] = 0
	}
}
