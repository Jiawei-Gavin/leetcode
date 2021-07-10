package leetcode

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	x := []int{1, 0, -1, 0}
	y := []int{0, 1, 0, -1}
	var queue [][]int
	queue = append(queue, entrance)
	maze[entrance[0]][entrance[1]] = '+'
	step := 0
	for len(queue) > 0 {
		cur := len(queue)
		for i := 0; i < cur; i++ {
			temp := queue[0]
			queue = queue[1:]
			row := temp[0]
			col := temp[1]
			if (row == 0 || col == 0 || row == m-1 || col == n-1) && !(row == entrance[0] && col == entrance[1]) {
				return step
			}
			for j := 0; j < 4; j++ {
				nRow := row + x[j]
				nCol := col + y[j]
				if nRow >= 0 && nCol >= 0 && nRow < m && nCol < n && maze[nRow][nCol] != '+' {
					maze[nRow][nCol] = '+'
					queue = append(queue, []int{nRow, nCol})
				}
			}
		}
		step++
	}
	return -1
}
