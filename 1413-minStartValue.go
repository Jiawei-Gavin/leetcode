package leetcode

func minStartValue(nums []int) int {
	sum, res := 0, 0
	for _, num := range nums {
		sum += num
		res = min(sum, res)
	}
	return 1 - res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
