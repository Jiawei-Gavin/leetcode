package leetcode

import "strings"

// solution1
func removeOccurrences(s string, part string) string {
	for strings.Contains(s, part) {
		s = strings.Replace(s, part, "", -1)
	}
	return s
}

func removeOccurrences(s string, part string) string {
	for {
		i := strings.Index(s, part)
		if i < 0 {
			break
		}
		s = s[:i] + s[i+len(part):]
	}
	return s
}
