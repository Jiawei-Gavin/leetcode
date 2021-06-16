package leetcode

func maximum(a int, b int) int {
	res := int64(a - b)
	res = int64(a) - res&(res>>63)
	return int(res)
}
