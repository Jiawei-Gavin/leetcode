package leetcode

func containsDuplicate(nums []int) bool {
	hashTable := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := hashTable[nums[i]]; ok {
			return true
		}
		hashTable[nums[i]] = i
	}
	return false
}
