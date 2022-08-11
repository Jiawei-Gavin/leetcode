package leetcode

func reformat(s string) string {
	i, j := 0, 1
	res := make([]byte, len(s)+1)
	for _, char := range s {
		if char < 58 {
			if i < len(s) {
				res[i] = byte(char)
				i += 2
			} else {
				return ""
			}
		} else {
			if j <= len(s) {
				res[j] = byte(char)
				j += 2
			} else {
				return ""
			}
		}
	}
	if res[len(s)] > 58 {
		res[len(s)-1] = res[0]
		return string(res[1:])
	}
	return string(res[:len(s)])
}
