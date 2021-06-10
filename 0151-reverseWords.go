package leetcode

import "strings"

func reverseWords(s string) string {
	a := strings.Fields(s)
	left, right := 0, len(a)-1
	for left < right {
		a[left], a[right] = a[right], a[left]
		left++
		right--
	}
	return strings.Join(a, " ")
}
