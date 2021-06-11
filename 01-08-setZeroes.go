package leetcode

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, col := make([]int, m), make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}
	for i := 0; i < m; i++ {
		if row[i] == 1 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 0; j < n; j++ {
		if col[j] == 1 {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}
