package leetcode

import "strings"

func repeatedStringMatch(a string, b string) int {
	lengthA, lengthB := len(a), len(b)
	maxLength := lengthA*2 + lengthB
	s, count := "", 0
	for len(s) < maxLength {
		if strings.Contains(s, b) {
			return count
		}
		s += a
		count++
	}
	return -1
}
