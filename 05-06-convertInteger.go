package leetcode

func convertInteger(A int, B int) int {
	res := 0
	for i, j := 0, A^B; i < 32; i++ {
		res += (j >> i) & 1
	}
	return res
}
