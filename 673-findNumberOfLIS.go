package leetcode

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	count := make([]int, len(nums))
	var max int
	for i := 0; i <= len(nums)-1; i += 1 {
		dp[i] = 1
		count[i] = 1
		for j := 0; j < i; j += 1 {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}

	var res int
	for i := 0; i <= len(dp)-1; i += 1 {
		if dp[i] == max {
			res += count[i]
		}
	}
	return res
}
