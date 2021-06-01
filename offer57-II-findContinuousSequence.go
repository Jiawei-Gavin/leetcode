package leetcode

func findContinuousSequence(target int) [][]int {
	left, right := 1, 2
	var result [][]int
	for left < right {
		sum := (left + right) * (right - left + 1) / 2
		if sum == target {
			var list []int
			for i := left; i <= right; i++ {
				list = append(list, i)
			}
			result = append(result, list)
			left++
		} else if sum < target {
			right++
		} else {
			left++
		}
	}
	return result
}
