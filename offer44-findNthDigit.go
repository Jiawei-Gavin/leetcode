package leetcode

import "strconv"

func findNthDigit(n int) int {
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count
		start *= 10
		digit++
		count = 9 * start * digit
	}
	num := start + (n-1)/digit
	return int(strconv.Itoa(num)[(n-1)%digit] - '0')
}
