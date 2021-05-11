package leetcode

var used [][]bool
var result bool

func exist(board [][]byte, word string) bool {
	result = false
	used = make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		used[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if word[0] == board[i][j] {
				search(board, word, 0, i, j)
				if result {
					return true
				}
			}
		}
	}
	return result
}

func search(board [][]byte, word string, idx, x, y int) {
	n := len(word)
	if result {
		return
	}
	if idx == n {
		result = true
		return
	}
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) ||
		used[x][y] ||
		word[idx] != board[x][y] {
		return
	}
	used[x][y] = true
	search(board, word, idx+1, x, y+1)
	search(board, word, idx+1, x+1, y)
	search(board, word, idx+1, x, y-1)
	search(board, word, idx+1, x-1, y)
	used[x][y] = false
}
