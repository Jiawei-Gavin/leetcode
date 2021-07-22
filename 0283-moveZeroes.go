package leetcode

func moveZeroes(nums []int) {
	for left, right := 0, 0; left < len(nums); left++ {
		if nums[left] != 0 {
			if left != right {
				nums[left], nums[right] = nums[right], nums[left]
			}
			right++
		}
	}
}
