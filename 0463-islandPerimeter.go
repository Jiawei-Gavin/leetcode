package leetcode

func islandPerimeter(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				sum := 4
				if i > 0 && grid[i-1][j] == 1 {
					sum--
				}
				if i < len(grid)-1 && grid[i+1][j] == 1 {
					sum--
				}
				if j > 0 && grid[i][j-1] == 1 {
					sum--
				}
				if j < len(grid[0])-1 && grid[i][j+1] == 1 {
					sum--
				}
				res += sum
			}
		}
	}
	return res
}
