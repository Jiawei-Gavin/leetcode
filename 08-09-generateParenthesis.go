package leetcode

func generateParenthesis(n int) []string {
	var res []string
	backtrack(0, 0, n, []byte{}, &res)
	return res
}

func backtrack(left, right, n int, path []byte, res *[]string) {
	if left > n || right > n || left < right {
		return
	}
	if len(path) == n*2 {
		*res = append(*res, string(path))
		return
	}
	backtrack(left+1, right, n, append(path, '('), res)
	backtrack(left, right+1, n, append(path, ')'), res)
}
