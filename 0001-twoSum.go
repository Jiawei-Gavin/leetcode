package leetcode

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := hashMap[another]; ok {
			return []int{hashMap[another], i}
		}
		hashMap[nums[i]] = i
	}
	return nil
}
