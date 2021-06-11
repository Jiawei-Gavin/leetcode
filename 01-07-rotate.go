package leetcode

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		l, r := 0, len(matrix)-1
		for l < r {
			matrix[l][i], matrix[r][i] = matrix[r][i], matrix[l][i]
			l++
			r--
		}
	}

	for i := 0; i < len(matrix)-1; i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
