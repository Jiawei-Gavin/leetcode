package leetcode

func add(a int, b int) int {
	for b != 0 {
		tmp := (a & b) << 1
		a ^= b
		b = tmp
	}
	return a
}
