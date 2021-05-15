package leetcode

func rotateTheBox(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])
	for row := 0; row < m; row++ {
		downEmptyPos := n - 1
		for col := n - 1; col >= 0; col-- {
			if box[row][col] == '#' {
				box[row][col] = '.'
				box[row][downEmptyPos] = '#'
				downEmptyPos -= 1
			} else if box[row][col] == '*' {
				downEmptyPos = col - 1
			}
		}
	}
	returnBox := make([][]byte, n)
	for i := 0; i < n; i++ {
		returnBox[i] = make([]byte, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			returnBox[i][j] = box[m-1-j][i]
		}
	}
	return returnBox
}
