package leetcode

func reverseParentheses(s string) string {
	pair := make([]int, len(s))
	var stack []int
	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else if v == ')' {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pair[i], pair[j] = j, i
		}
	}

	var result []byte
	for i, j := 0, 1; i < len(s); i += j {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			j = -j
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}
