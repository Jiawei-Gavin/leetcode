package leetcode

func smallestDivisor(nums []int, threshold int) int {
	maxNum := 1
	for _, num := range nums {
		maxNum = max(maxNum, num)
	}
	left, right := 1, maxNum
	for left < right {
		mid := (left + right) >> 1
		if calculateSum(nums, mid) > threshold {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func calculateSum(nums []int, divisor int) int {
	sum := 0
	for _, num := range nums {
		sum += num / divisor
		if num%divisor != 0 {
			sum++
		}
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
