package leetcode

import "strconv"

func printNumbers(n int) []int {
	var result []int
	str := make([]byte, n)
	dfs(0, n, str, &result)
	return result
}
func dfs(dep int, n int, str []byte, result *[]int) {
	if dep == n {
		val, _ := strconv.Atoi(string(str))
		if val != 0 {
			*result = append(*result, val)
		}
		return
	}
	for i := 0; i < 10; i++ {
		str[dep] = byte('0' + i)
		dfs(dep+1, n, str, result)
	}
}
