package leetcode

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	result := 0
	for i, num := range nums[:n/2] {
		result = max(result, num+nums[n-1-i])
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
