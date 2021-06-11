package leetcode

import (
	"bytes"
	"strings"
)

// solution1
func replaceSpaces(S string, length int) string {
	return strings.Replace(S[:length], " ", "%20", -1)
}

// solution2
func replaceSpaces(S string, length int) string {
	var res []byte
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			res = append(res, []byte{'%', '2', '0'}...)
		} else {
			res = append(res, S[i])
		}
	}
	return string(res)
}
