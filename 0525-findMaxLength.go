package leetcode

func findMaxLength(nums []int) int {
	hashMap := map[int]int{0: -1}
	result, count := 0, 0
	for i, num := range nums {
		if num == 1 {
			count++
		} else {
			count--
		}
		if index, ok := hashMap[count]; ok {
			result = max(result, i-index)
		} else {
			hashMap[count] = i
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
