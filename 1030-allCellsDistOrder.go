package leetcode

import "sort"

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	res := make([][]int, 0, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res = append(res, []int{i, j})
		}
	}
	sort.Slice(res, func(i, j int) bool {
		a, b := res[i], res[j]
		return abs(a[0]-rCenter)+abs(a[1]-cCenter) < abs(b[0]-rCenter)+abs(b[1]-cCenter)
	})
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
