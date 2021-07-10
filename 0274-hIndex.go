package leetcode

import "sort"

// solution1
func hIndex(citations []int) int {
	sort.Ints(citations)
	h := 0
	for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}
	return h
}

// solution2
func hIndex(citations []int) int {
	left, right := 0, len(citations)
	for left <= right {
		mid := (left + right) >> 1
		count := 0
		for _, citation := range citations {
			if citation >= mid {
				count++
			}
		}
		if count < mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
