package leetcode

import "strings"

func reverseWords(s string) string {
	a := strings.Split(s, " ")
	for i, v := range a {
		a[i] = reverse(v)
	}
	return strings.Join(a, " ")
}

func reverse(s string) string {
	bytes := []byte(s)
	left, right := 0, len(bytes)-1
	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}
	return string(bytes)
}
