package leetcode

func titleToNumber(columnTitle string) int {
	val, result := 0, 0
	for i := 0; i < len(columnTitle); i++ {
		val = int(columnTitle[i] - 'A' + 1)
		result = result*26 + val
	}
	return result
}
