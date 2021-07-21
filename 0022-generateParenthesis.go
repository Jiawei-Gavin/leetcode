package leetcode

func generateParenthesis(n int) []string {
	var res []string
	dfs(n, n, "", &res)
	return res
}

func dfs(left, right int, tmp string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, tmp)
		return
	}
	if left > right {
		return
	}
	if left > 0 {
		dfs(left-1, right, tmp+"(", res)
	}
	if right > 0 {
		dfs(left, right-1, tmp+")", res)
	}
}
