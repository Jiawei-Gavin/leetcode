package leetcode

func xorQueries(arr []int, queries [][]int) []int {
	pre := make([]int, len(arr)+1)
	pre[0] = 0
	for i, val := range arr {
		pre[i+1] = pre[i] ^ val
	}

	result := make([]int, len(queries))
	for i := 0; i < len(result); i++ {
		result[i] = pre[queries[i][0]] ^ pre[queries[i][1]+1]
	}
	return result
}
