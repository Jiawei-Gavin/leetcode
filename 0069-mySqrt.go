package leetcode

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right, result := 1, x, 0
	for left <= right {
		mid := left + ((right - left) >> 1)
		if mid < x/mid {
			left = mid + 1
			result = mid
		} else if mid == x/mid {
			return mid
		} else {
			right = mid - 1
		}
	}
	return result
}
