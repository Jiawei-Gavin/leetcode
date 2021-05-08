package leetcode

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	stack := make([]byte, 0)
	for i := 0; i < n; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if ((s[i] == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((s[i] == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') ||
			((s[i] == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
