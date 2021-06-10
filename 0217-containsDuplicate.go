package leetcode

func containsDuplicate(nums []int) bool {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := hashMap[nums[i]]; ok {
			return true
		}
		hashMap[nums[i]] = i
	}
	return false
}
