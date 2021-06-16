package leetcode

import (
	"math"
	"sort"
)

func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	i, j, res := 0, 0, math.MaxInt32
	for i < len(a) && j < len(b) {
		res = min(res, abs(a[i], b[j]))
		if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}
