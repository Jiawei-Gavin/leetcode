package leetcode

import "sort"

func minSubsequence(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	total := 0
	for _, num := range nums {
		total += num
	}
	for i, sum := 0, 0; ; i++ {
		sum += nums[i]
		if sum > total-sum {
			return nums[:i+1]
		}
	}
}
