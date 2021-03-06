package leetcode

func numberOf2sInRange(n int) int {
	digit, res := 1, 0
	high, cur, low := n/10, n%10, 0
	for high != 0 || cur != 0 {
		if cur < 2 {
			res += high * digit
		} else if cur == 2 {
			res += high*digit + low + 1
		} else {
			res += (high + 1) * digit
		}
		low += cur * digit
		high, cur = high/10, high%10
		digit = digit * 10
	}
	return res
}
