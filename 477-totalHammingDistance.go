package leetcode

func totalHammingDistance(nums []int) int {
	result := 0
	for i := 0; i < 30; i++ {
		c := 0
		for _, num := range nums {
			c += num >> i & 1
		}
		result += c * (len(nums) - c)
	}
	return result
}
