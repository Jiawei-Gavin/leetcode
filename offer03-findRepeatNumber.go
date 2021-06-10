package leetcode

func findRepeatNumber(nums []int) int {
	hashMap := make(map[int]bool)
	for _, num := range nums {
		if !hashMap[num] {
			hashMap[num] = true
		} else {
			return num
		}
	}
	return -1
}
