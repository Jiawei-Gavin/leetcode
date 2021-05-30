package leetcode

func findRepeatNumber(nums []int) int {
	hashTable := make(map[int]bool)
	for _, num := range nums {
		if !hashTable[num] {
			hashTable[num] = true
		} else {
			return num
		}
	}
	return -1
}
