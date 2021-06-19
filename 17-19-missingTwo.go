package leetcode

func missingTwo(nums []int) []int {
	nums = append(nums, 0, -1, -1)
	res := make([]int, 0, 2)
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] != i && nums[i] != -1 {
			j := nums[i]
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	for i, num := range nums {
		if num == -1 {
			res = append(res, i)
		}
	}
	return res
}
