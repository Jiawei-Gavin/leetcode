package leetcode

func lenLongestFibSubseq(arr []int) int {
	var res int
	if len(arr) < 3 {
		return res
	}
	dp := make([][]int, len(arr)-1)
	for i := range dp {
		for j := 0; j < len(arr); j++ {
			dp[i] = append(dp[i], 2)
		}
	}
	index := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		index[arr[i]] = i
		for j := 0; j < i; j++ {
			tmp := arr[i] - arr[j]
			if k, ok := index[tmp]; ok == true && k < j {
				dp[j][i] = max(dp[j][i], dp[k][j]+1)
			}
			res = max(res, dp[j][i])
		}
	}
	if res == 2 {
		res = 0
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
