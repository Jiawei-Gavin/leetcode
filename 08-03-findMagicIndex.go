package leetcode

func findMagicIndex(nums []int) int {
	for i, num := range nums {
		if i == num {
			return i
		}
	}
	return -1
}
