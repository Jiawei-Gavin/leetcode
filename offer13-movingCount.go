package leetcode

var m1, n1, k1 int
var visited [][]bool

func movingCount(m int, n int, k int) int {
	m1, n1, k1 = m, n, k
	visited = make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}
	return dfs(0, 0, 0, 0)
}

func dfs(i, j, si, sj int) int {
	if i >= m1 || j >= n1 || si+sj > k1 || visited[i][j] {
		return 0
	}
	visited[i][j] = true
	si1 := si + 1
	sj1 := sj + 1
	if (i+1)%10 == 0 {
		si1 = si - 8
	}
	if (j+1)%10 == 0 {
		sj1 = sj - 8
	}
	return 1 + dfs(i, j+1, si, sj1) + dfs(i+1, j, si1, sj)
}
