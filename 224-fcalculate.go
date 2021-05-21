package leetcode

import "container/list"

func calculate(s string) int {
	stack, sign, i, result := list.New(), 1, 0, 0
	for i < len(s) {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = 1
			i++
		case '-':
			sign = -1
			i++
		case '(':
			stack.PushBack(result)
			stack.PushBack(sign)
			result = 0
			sign = 1
			i++
		case ')':
			result = result*stack.Remove(stack.Back()).(int) + stack.Remove(stack.Back()).(int)
			i++
		default:
			num := 0
			for i < len(s) && s[i] <= '9' && s[i] >= '0' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			result += num * sign
		}
	}
	return result
}
