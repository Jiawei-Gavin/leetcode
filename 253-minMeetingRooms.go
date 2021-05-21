package leetcode

import "sort"

func minMeetingRooms(intervals [][]int) int {
	var nums []int
	for _, interval := range intervals {
		nums = append(nums, interval[0]*10+2)
		nums = append(nums, interval[1]*10+1)
	}
	sort.Ints(nums)
	result, cur := 0, 0
	for _, v := range nums {
		if v%10 == 1 {
			cur--
		} else {
			cur++
		}
		if cur > result {
			result = cur
		}
	}
	return result
}
