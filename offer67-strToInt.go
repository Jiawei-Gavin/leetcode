package leetcode

func strToInt(s string) int {
	isNegative := false
	result := 0
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && j == 0 {
			continue
		}
		if s[i] == '-' && j == 0 {
			isNegative = true
			j++
			continue
		}
		if s[i] == '+' && j == 0 {
			isNegative = false
			j++
			continue
		}
		if !isNum(s[i]) && j == 0 {
			return 0
		}
		if !isNum(s[i]) && j > 0 {
			break
		}
		j++
		if result*10+int(s[i]-'0') > 1<<31-1 {
			if isNegative {
				return -(1 << 31)
			}
			return 1<<31 - 1
		}
		result = result*10 + int(s[i]-'0')
	}
	if isNegative {
		result = -result
	}
	return result
}

func isNum(c byte) bool {
	return c >= '0' && c <= '9'
}
