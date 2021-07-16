package leetcode

func strWithout3a3b(a int, b int) string {
	var res string
	for a > 0 && b > 0 {
		if a > b {
			res += "aab"
			a -= 2
			b -= 1
		} else if a == b {
			res += "ab"
			a -= 1
			b -= 1
		} else {
			res += "bba"
			a -= 1
			b -= 2
		}
	}
	for a > 0 {
		res += "a"
		a -= 1
	}
	for b > 0 {
		res += "b"
		b -= 1
	}
	return res
}
