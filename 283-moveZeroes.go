package leetcode

// solution1
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

// solution2
func moveZeroes(nums []int) {
	right := 0
	for left := 0; left < len(nums); left++ {
		if nums[left] != 0 {
			if left != right {
				nums[left], nums[right] = nums[right], nums[left]
			}
			right++
		}
	}
}
