package leetcode

func missingNumber(nums []int) int {
	var res int
	for i := 0; i <= len(nums); i++ {
		if i == len(nums) {
			res ^= i
		} else {
			res ^= i ^ nums[i]
		}
	}
	return res
}
