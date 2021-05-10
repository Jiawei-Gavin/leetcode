package leetcode

// solution1
func numIslands(grid [][]byte) int {
	result := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				result++
				dfs(i, j, grid)
			}
		}
	}
	return result
}

func dfs(x, y int, grid [][]byte) {
	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == '1' {
		grid[x][y] = '0'
		dfs(x+1, y, grid)
		dfs(x-1, y, grid)
		dfs(x, y+1, grid)
		dfs(x, y-1, grid)
	}
}

// solution2
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	var result int
	var queue [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				result++
				grid[i][j] = '0'
				queue = append(queue, []int{i, j})
				for len(queue) != 0 {
					cur := queue[len(queue)-1]
					queue = queue[:len(queue)-1]
					x := cur[0]
					y := cur[1]
					if x-1 >= 0 && grid[x-1][y] == '1' {
						grid[x-1][y] = '0'
						queue = append(queue, []int{x - 1, y})
					}
					if y-1 >= 0 && grid[x][y-1] == '1' {
						grid[x][y-1] = '0'
						queue = append(queue, []int{x, y - 1})
					}
					if x+1 < m && grid[x+1][y] == '1' {
						grid[x+1][y] = '0'
						queue = append(queue, []int{x + 1, y})
					}
					if y+1 < n && grid[x][y+1] == '1' {
						grid[x][y+1] = '0'
						queue = append(queue, []int{x, y + 1})
					}
				}
			}
		}
	}
	return result
}
