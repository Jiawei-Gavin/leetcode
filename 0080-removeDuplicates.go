package leetcode

func removeDuplicates(nums []int) int {
	res := 0
	for _, num := range nums {
		if res < 2 || nums[res-2] != num {
			nums[res] = num
			res++
		}
	}
	return res
}
