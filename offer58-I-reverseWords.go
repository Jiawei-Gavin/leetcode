package leetcode

import "strings"

func reverseWords(s string) string {
	str := strings.Fields(s)
	left, right := 0, len(str)-1
	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
	return strings.Join(str, " ")
}
