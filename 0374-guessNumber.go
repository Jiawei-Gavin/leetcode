package leetcode

func guessNumber(n int) int {
	left, right := 1, n
	for left < right {
		mid := (left + right) / 2
		if guess(mid) == 1 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
