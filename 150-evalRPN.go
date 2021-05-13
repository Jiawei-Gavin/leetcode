package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	if len(tokens) == 1 {
		i, _ := strconv.Atoi(tokens[0])
		return i
	}
	top := 0
	for _, token := range tokens {
		switch token {
		case "+":
			{
				sum := stack[top-2] + stack[top-1]
				stack = stack[:top-2]
				stack = append(stack, sum)
				top--
			}
		case "-":
			{
				sub := stack[top-2] - stack[top-1]
				stack = stack[:top-2]
				stack = append(stack, sub)
				top--
			}
		case "*":
			{
				mul := stack[top-2] * stack[top-1]
				stack = stack[:top-2]
				stack = append(stack, mul)
				top--
			}
		case "/":
			{
				div := stack[top-2] / stack[top-1]
				stack = stack[:top-2]
				stack = append(stack, div)
				top--
			}
		default:
			{
				i, _ := strconv.Atoi(token)
				stack = append(stack, i)
				top++
			}
		}
	}
	return stack[0]
}
