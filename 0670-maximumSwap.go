package leetcode

import "strconv"

func maximumSwap(num int) int {
	s := strconv.Itoa(num)
	nums := []byte(s)
	arr := make([]int, len(nums))
	end := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[end] {
			end = i
		}
		arr[i] = end
	}
	for i := range nums {
		if nums[arr[i]] != nums[i] {
			nums[i], nums[arr[i]] = nums[arr[i]], nums[i]
			break
		}
	}
	res, _ := strconv.Atoi(string(nums))
	return res
}
