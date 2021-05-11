package leetcode

func generateParenthesis(n int) []string {
	var result []string
	parenthesisFunc(n, n, "", &result)
	return result
}

func parenthesisFunc(left, right int, tmp string, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, tmp)
		return
	}
	if left > 0 {
		parenthesisFunc(left-1, right, tmp+"(", result)
	}
	if right > left {
		parenthesisFunc(left, right-1, tmp+")", result)
	}
}
