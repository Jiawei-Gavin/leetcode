package leetcode

func multiply(A int, B int) int {
	if A == 1 {
		return B
	}
	if B == 1 {
		return A
	}
	if B&1 == 1 {
		return A + multiply(A<<1, B>>1)
	} else {
		return multiply(A<<1, B>>1)
	}
}
