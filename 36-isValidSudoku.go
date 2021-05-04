package leetcode

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		row := map[byte]bool{}
		column := map[byte]bool{}
		cell := map[byte]bool{}
		for j := 0; j < 9; j++ {
			// row
			if board[i][j] != '.' {
				if row[board[i][j]] {
					return false
				}
				row[board[i][j]] = true
			}
			// column
			if board[j][i] != '.' {
				if column[board[j][i]] {
					return false
				}
				column[board[j][i]] = true
			}
			// 3*3 cell
			cellRow := (i%3)*3 + j%3
			cellCol := (i/3)*3 + j/3
			if board[cellRow][cellCol] != '.' {
				if cell[board[cellRow][cellCol]] {
					return false
				}
				cell[board[cellRow][cellCol]] = true
			}
		}
	}
	return true
}
