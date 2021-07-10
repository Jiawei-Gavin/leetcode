package leetcode

import "sort"

func findRightInterval(intervals [][]int) []int {
	res := make([]int, len(intervals))
	hashMap := make(map[int]int)
	for i, interval := range intervals {
		hashMap[interval[0]] = i
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := range intervals {
		l, r := i, len(intervals)-1
		for l < r {
			mid := (l + r) >> 1
			if intervals[mid][0] >= intervals[i][1] {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if intervals[i][1] > intervals[r][0] || i == len(intervals)-1 {
			res[hashMap[intervals[i][0]]] = -1
		} else {
			res[hashMap[intervals[i][0]]] = hashMap[intervals[r][0]]
		}
	}
	return res
}
