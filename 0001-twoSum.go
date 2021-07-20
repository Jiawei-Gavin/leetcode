package leetcode

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		if _, ok := hashMap[temp]; ok {
			return []int{hashMap[temp], i}
		}
		hashMap[nums[i]] = i
	}
	return nil
}
