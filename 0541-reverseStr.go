package leetcode

func reverseStr(s string, k int) string {
	b := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		if i+k <= len(s) {
			reverse(b[i : i+k])
		} else {
			reverse(b[i:])
		}
	}
	return string(b)
}

func reverse(b []byte) {
	start, end := 0, len(b)-1
	for start < end {
		b[start], b[end] = b[end], b[start]
	}
}
