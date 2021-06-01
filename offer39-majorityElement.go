package leetcode

func majorityElement(nums []int) int {
	result, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			result = num
			count++
		} else if result != num {
			count--
		} else {
			count++
		}
	}
	return result
}
