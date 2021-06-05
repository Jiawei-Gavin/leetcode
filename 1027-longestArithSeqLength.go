package leetcode

func longestArithSeqLength(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	dp := make([]map[int]int, len(nums)+1)
	var res int
	for i := 0; i < len(dp); i++ {
		dp[i] = make(map[int]int)
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			tmp := nums[i] - nums[j]
			if _, ok := dp[j][tmp]; ok {
				dp[i][tmp] = dp[j][tmp] + 1
			} else {
				dp[j][tmp] = 1
				dp[i][tmp] = dp[j][tmp] + 1
			}
			res = max(res, dp[i][tmp])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
