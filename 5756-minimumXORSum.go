package leetcode

import "math"

func minimumXORSum(nums1 []int, nums2 []int) int {
	n := len(nums1)
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := (1 << n) - 1; j >= 0; j-- {
			if dp[j] < math.MaxInt32 {
				for k := 0; k < n; k++ {
					if (j & (1 << k)) == 0 {
						nxt := j ^ (1 << k)
						dp[nxt] = min(dp[nxt], dp[j]+(nums1[i]^nums2[k]))
					}
				}
			}
		}
	}
	return dp[(1<<n)-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
