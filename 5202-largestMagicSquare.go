package leetcode

func largestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rs := make([][]int, m)
	cs := make([][]int, m+1)
	ds := make([][]int, m+1)
	as := make([][]int, m+1)
	for i := range cs {
		cs[i] = make([]int, n)
		ds[i] = make([]int, n+1)
		as[i] = make([]int, n+1)
	}
	for i, row := range grid {
		rs[i] = make([]int, n+1)
		for j, v := range row {
			rs[i][j+1] = rs[i][j] + v
			cs[i+1][j] = cs[i][j] + v
			ds[i+1][j+1] = ds[i][j] + v
			as[i+1][j] = as[i][j+1] + v
		}
	}

	for k := min(m, n); ; k-- {
		for r := k; r <= m; r++ {
		Next:
			for c := k; c <= n; c++ {
				s := ds[r][c] - ds[r-k][c-k]
				if as[r][c-k]-as[r-k][c] != s {
					continue
				}
				for _, row := range rs[r-k : r] {
					if row[c]-row[c-k] != s {
						continue Next
					}
				}
				for j := c - k; j < c; j++ {
					if cs[r][j]-cs[r-k][j] != s {
						continue Next
					}
				}
				return k
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
