package leetcode

func decode(encoded []int, first int) []int {
	result := make([]int, len(encoded)+1)
	result[0] = first
	for i := 0; i < len(encoded); i++ {
		result[i+1] = encoded[i] ^ result[i]
	}
	return result
}
