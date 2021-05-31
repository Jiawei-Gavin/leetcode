package leetcode

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left] < numbers[right] {
			return numbers[left]
		}
		mid := (left + right) / 2
		if numbers[left] < numbers[mid] {
			left = mid + 1
		} else if numbers[mid] == numbers[left] {
			left++
		} else {
			right = mid
		}
	}
	return numbers[left]
}
