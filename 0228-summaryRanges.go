package leetcode

import "strconv"

func summaryRanges(nums []int) []string {
	var res []string
	left, right := 0, 0
	for left < len(nums) {
		for right < len(nums)-1 && nums[right+1]-nums[right] == 1 {
			right++
		}
		temp := strconv.Itoa(nums[left])
		if left == right {
			res = append(res, temp)
		} else {
			res = append(res, temp+"->"+strconv.Itoa(nums[right]))
		}
		right++
		left = right
	}
	return res
}
