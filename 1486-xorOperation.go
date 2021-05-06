package leetcode

// solution1
func xorOperation(n int, start int) int {
	result := 0
	for i := 0; i < n; i++ {
		result ^= start + 2*i
	}
	return result
}

// solution2 ---- 数学思维很重要！
func xorOperation(n int, start int) int {
	s, e := start>>1, n&start&1
	result := sumXor(s-1) ^ sumXor(s+n-1)
	return result<<1 | e
}

func sumXor(x int) int {
	switch x % 4 {
	case 0:
		return x
	case 1:
		return 1
	case 2:
		return x + 1
	default:
		return 0
	}
}
