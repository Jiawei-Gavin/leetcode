package leetcode

func grayCode(n int) []int {
	res := make([]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		res[i] = i ^ i>>1
	}
	return res
}
