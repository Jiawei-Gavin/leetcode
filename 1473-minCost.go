package leetcode

import "math"

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	dp := make([][][]int, m)
	for i, _ := range dp {
		dp[i] = make([][]int, m+1)
		for j, _ := range dp[i] {
			dp[i][j] = make([]int, n)
			for k, _ := range dp[i][j] {
				dp[i][j][k] = math.MaxInt32
			}
		}
	}

	if houses[0] != 0 {
		dp[0][1][houses[0]-1] = 0
	} else {
		for i := 0; i < n; i++ {
			dp[0][1][i] = cost[0][i]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < target+1; j++ {
			for k := 1; k < n+1; k++ {
				temp := math.MaxInt32
				for g := 1; g < n+1; g++ {
					if g != k {
						temp = min(temp, dp[i-1][j-1][g-1])
					}
				}
				temp = min(temp, dp[i-1][j][k-1])
				if houses[i] == 0 {
					dp[i][j][k-1] = temp + cost[i][k-1]
				} else if houses[i] == k {
					dp[i][j][k-1] = temp
				}
			}
		}
	}
	res := math.MaxInt32
	for i := 0; i < n; i++ {
		res = min(dp[m-1][target][i], res)
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
