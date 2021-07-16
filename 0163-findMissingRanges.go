package leetcode

import "strconv"

func findMissingRanges(nums []int, lower int, upper int) []string {
	var res []string
	pre := lower - 1
	for _, num := range nums {
		if num-pre == 2 {
			res = append(res, strconv.Itoa(pre+1))
		} else if num-pre > 2 {
			temp := strconv.Itoa(pre+1) + "->" + strconv.Itoa(num-1)
			res = append(res, temp)
		}
		pre = num
	}
	if upper-pre == 1 {
		res = append(res, strconv.Itoa(pre+1))
	} else if upper-pre > 1 {
		temp := strconv.Itoa(pre+1) + "->" + strconv.Itoa(upper)
		res = append(res, temp)
	}
	return res
}
