package leetcode

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := hashTable[another]; ok {
			return []int{hashTable[another], i}
		}
		hashTable[nums[i]] = i
	}
	return nil
}
