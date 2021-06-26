package leetcode

import "math"

func maxAlternatingSum(nums []int) int64 {
	f := [2]int{0, math.MinInt64}
	for _, num := range nums {
		f = [2]int{max(f[0], f[1]-num), max(max(f[1], f[0]+num), num)}
	}
	return int64(f[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
