package leetcode

func replaceSpace(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			result += "%20"
		} else {
			result += string(s[i])
		}
	}
	return result
}
