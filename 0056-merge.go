package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var result [][]int
	pre := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if pre[1] < cur[0] {
			result = append(result, pre)
			pre = cur
		} else {
			pre[1] = max(pre[1], cur[1])
		}
	}
	result = append(result, pre)
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
