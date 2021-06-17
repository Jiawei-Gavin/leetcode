package leetcode

func calculate(s string) int {
	var stack []int
	sign, num, res := byte('+'), 0, 0
	for i := 0; i < len(s); i++ {
		isDigit := s[i] <= '9' && s[i] >= '0'
		if isDigit {
			num = num*10 + int(s[i]-'0')
		}
		if !isDigit && s[i] != ' ' || i == len(s)-1 {
			if sign == '+' {
				stack = append(stack, num)
			} else if sign == '-' {
				stack = append(stack, -num)
			} else if sign == '*' {
				stack[len(stack)-1] *= num
			} else if sign == '/' {
				stack[len(stack)-1] /= num
			}
			sign = s[i]
			num = 0
		}
	}
	for _, i := range stack {
		res += i
	}
	return res
}
