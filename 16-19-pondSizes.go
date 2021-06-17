package leetcode

import "sort"

func pondSizes(land [][]int) []int {
	var res []int
	for i := range land {
		for j := range land[i] {
			tmp := dfs(&land, i, j)
			if tmp > 0 {
				res = append(res, tmp)
			}
		}
	}
	sort.Ints(res)
	return res
}

func dfs(land *[][]int, x, y int) int {
	if x < 0 || y < 0 || x >= len(*land) || y >= len((*land)[0]) {
		return 0
	}

	if (*land)[x][y] == 0 {
		(*land)[x][y] = -1
		return 1 + dfs(land, x, y+1) + dfs(land, x+1, y) + dfs(land, x-1, y) + dfs(land, x, y-1) + dfs(land, x-1, y-1) + dfs(land, x-1, y+1) + dfs(land, x, y-1) + dfs(land, x, y+1) + dfs(land, x+1, y-1) + dfs(land, x+1, y+1)
	}
	return 0
}
