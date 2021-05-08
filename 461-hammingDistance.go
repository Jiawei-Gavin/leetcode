package leetcode

import "math/bits"

// solution1
func hammingDistance(x int, y int) int {
	xor := x ^ y
	return bits.OnesCount(uint(xor))
}

// solution2
func hammingDistance(x int, y int) int {
	distance := 0
	for xor := x ^ y; xor != 0; xor &= xor - 1 {
		distance++
	}
	return distance
}
