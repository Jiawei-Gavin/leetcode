package leetcode

import "sort"

func subSort(array []int) []int {
	sorted := make([]int, len(array))
	copy(sorted, array)
	sort.Ints(sorted)
	left, right := -1, -1
	for i := 0; i < len(array); i++ {
		if sorted[i] != array[i] {
			left = i
			break
		}
	}
	for i := len(array) - 1; i > 0; i-- {
		if sorted[i] != array[i] {
			right = i
			break
		}
	}
	return []int{left, right}
}
