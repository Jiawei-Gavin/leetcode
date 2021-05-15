package leetcode

func findMaximumXOR(nums []int) int {
	result, mask := 0, 0
	for i := 31; i >= 0; i-- {
		mask = mask | (1 << uint(i))
		m := make(map[int]bool)
		for _, num := range nums {
			m[num&mask] = true
		}
		greedyTry := result | (1 << uint(i))
		for anotherNum := range m {
			if m[anotherNum^greedyTry] == true {
				result = greedyTry
				break
			}
		}
	}
	return result
}
