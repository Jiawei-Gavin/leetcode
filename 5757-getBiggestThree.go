package leetcode

import "sort"

func getBiggestThree(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	var result []int
	for i := range grid {
		for j := range grid[0] {
			result = append(result, grid[i][j])
			length := min(min(i, m-1-i), min(j, n-1-j))
			for l := 1; l < length+1; l++ {
				cur := 0
				cur += grid[i][j-l]
				cur += grid[i][j+l]
				cur += grid[i-l][j]
				cur += grid[i+l][j]
				for k := 1; k < l; k++ {
					cur += grid[i-k][j-l+k]
					cur += grid[i+k][j-l+k]
					cur += grid[i-k][j+l-k]
					cur += grid[i+k][j+l-k]
				}
				result = append(result, cur)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	result = removeDuplicate(result)
	if len(result) <= 3 {
		return result
	}
	return result[:3]
}

func removeDuplicate(res []int) []int {
	hashTable := map[int]bool{}
	var result []int
	for _, v := range res {
		if !hashTable[v] {
			hashTable[v] = true
			result = append(result, v)
		}
	}
	return result
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
