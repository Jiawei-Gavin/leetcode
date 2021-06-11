package leetcode

import (
	"strconv"
	"strings"
)

func compressString(S string) string {
	if S == "" {
		return ""
	}
	var str strings.Builder
	i := 0
	for i < len(S) {
		count := 1
		for j := i + 1; j < len(S); j++ {
			if S[j] == S[i] {
				count++
			} else {
				str.WriteString(string(S[i]))
				str.WriteString(strconv.Itoa(count))
				break
			}
		}
		i = i + count
		if i == len(S) {
			str.WriteString(string(S[i-count]))
			str.WriteString(strconv.Itoa(count))
		}
	}
	res := str.String()
	if len(res) >= len(S) {
		return S
	}
	return res
}
