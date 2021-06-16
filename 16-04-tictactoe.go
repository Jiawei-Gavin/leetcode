package leetcode

func tictactoe(board []string) string {
	pending := 0
	hashMap := map[string]int{
		"X": 1,
		"O": 0,
		" ": -3,
	}

	for i := 0; i < len(board); i++ {
		sum := 0
		for j := 0; j < len(board); j++ {
			sum += hashMap[string(board[j][i])]
		}
		if sum == len(board) {
			return "X"
		} else if sum == 0 {
			return "O"
		} else if sum < 0 {
			pending = 1
		}
	}

	for i := 0; i < len(board); i++ {
		sum := 0
		for j := 0; j < len(board); j++ {
			sum += hashMap[string(board[i][j])]
		}
		if sum == len(board) {
			return "X"
		} else if sum == 0 {
			return "O"
		} else if sum < 0 {
			pending = 1
		}
	}

	sum1, sum2 := 0, 0
	for i := 0; i < len(board); i++ {
		sum1 = sum1 + hashMap[string(board[len(board)-i-1][i])]
		sum2 = sum2 + hashMap[string(board[i][i])]
	}

	if sum1 == len(board) || sum2 == len(board) {
		return "X"
	} else if sum1 == 0 || sum2 == 0 {
		return "O"
	} else if sum1 < 0 || sum2 < 0 {
		pending = 1
	}

	if pending == 1 {
		return "Pending"
	} else {
		return "Draw"
	}
}
