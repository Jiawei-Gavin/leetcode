package leetcode

import "math/bits"

// solution1
func hammingWeight(num uint32) int {
	return bits.OnesCount(uint(num))
}

// solution2
func hammingWeight(num uint32) int {
	ones := 0
	for num != 0 {
		num = num & (num - 1)
		ones++
	}
	return ones
}
