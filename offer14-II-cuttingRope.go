package leetcode

func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	result := 1
	for n > 4 {
		result = result * 3 % 1000000007
		n -= 3
	}
	return result * n % 1000000007
}
