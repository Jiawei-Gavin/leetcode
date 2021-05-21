package leetcode

func xorGame(nums []int) bool {
	if len(nums)%2 == 0 {
		return true
	}
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result == 0
}
