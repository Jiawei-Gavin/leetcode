package leetcode

import "strings"

func addBoldTag(s string, words []string) string {
	mask := make([]bool, len(s))
	end := 0
	for i := 0; i < len(s); i++ {
		for _, word := range words {
			if i+len(word) <= len(s) && s[i:i+len(word)] == word {
				end = max(end, i+len(word))
			}
		}
		mask[i] = end > i
	}
	var res strings.Builder
	for i := 0; i < len(s); i++ {
		if mask[i] && (i == 0 || !mask[i-1]) {
			res.WriteString("<b>")
		}
		res.WriteByte(s[i])
		if mask[i] && (i == len(s)-1 || !mask[i+1]) {
			res.WriteString("</b>")
		}
	}
	return res.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
