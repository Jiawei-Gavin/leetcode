package leetcode

import "sort"

// solution1
func missingNumber(nums []int) int {
	sort.Ints(nums)
	start, end := 0, len(nums)
	for start < end {
		mid := start + (end-start)>>1
		if nums[mid] > mid {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

// solution2
func missingNumber(nums []int) int {
	xor, i := 0, 0
	for i = 0; i < len(nums); i++ {
		xor = xor ^ i ^ nums[i]
	}
	return xor ^ i
}
