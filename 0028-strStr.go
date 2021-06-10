package leetcode

func strStr(haystack string, needle string) int {
next:
	for i := 0; i+len(needle) <= len(haystack); i++ {
		for j := range needle {
			if needle[j] != haystack[i+j] {
				continue next
			}
		}
		return i
	}
	return -1
}
