package leetcode

func arrangeCoins(n int) int {
	left, right := 0, n/2+1
	for left < right {
		mid := left + (right-left+1)>>1
		val := mid * (mid + 1) / 2
		if val > n {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
