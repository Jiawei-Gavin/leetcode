package leetcode

// solution1
func majorityElement(nums []int) int {
	hashMap := make(map[int]int)
	for _, val := range nums {
		hashMap[val]++
		if hashMap[val] > len(nums)/2 {
			return val
		}
	}
	return 0
}

// solution2
func majorityElement(nums []int) int {
	result, count := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			result, count = nums[i], 1
		} else {
			if nums[i] == result {
				count++
			} else {
				count--
			}
		}
	}
	return result
}
