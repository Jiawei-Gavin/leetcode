package leetcode

func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for _, v := range chalk {
		sum += v
	}
	k %= sum
	for i, v := range chalk {
		if v > k {
			return i
		}
		k -= v
	}
	return 0
}
