package leetcode

import "strings"

// solution1
func countSegments(s string) int {
	res := 0
	for i, v := range s {
		if (i == 0 || s[i-1] == ' ') && v != ' ' {
			res++
		}
	}
	return res
}

// solution2
func countSegments(s string) int {
	return len(strings.Fields(s))
}
