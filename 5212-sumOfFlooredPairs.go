package leetcode

func sumOfFlooredPairs(nums []int) int {
	result, cnt := 0, make([]int, 2e5)
	for _, v := range nums {
		cnt[v]++
	}
	sum := make([]int, 2e5+1)
	for i, v := range cnt {
		sum[i+1] = sum[i] + v
	}
	for i, v := range cnt {
		if v > 0 {
			for j := 1; j*i <= 1e5; j++ {
				result += v * j * (sum[(j+1)*i] - sum[j*i])
			}
		}
	}
	return result % (1e9 + 7)
}
