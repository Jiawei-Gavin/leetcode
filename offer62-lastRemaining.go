package leetcode

func lastRemaining(n int, m int) int {
	var res int
	for i := 2; i <= n; i++ {
		res = (res + m) % i
	}
	return res
}
