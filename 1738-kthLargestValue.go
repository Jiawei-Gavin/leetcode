package leetcode

import "sort"

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)
	pre := make([][]int, m+1)
	pre[0] = make([]int, n+1)
	for i, row := range matrix {
		pre[i+1] = make([]int, n+1)
		for j, v := range row {
			pre[i+1][j+1] = pre[i+1][j] ^ pre[i][j+1] ^ pre[i][j] ^ v
			result = append(result, pre[i+1][j+1])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result[k-1]
}
