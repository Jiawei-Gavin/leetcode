package leetcode

func toLowerCase(s string) string {
	var res string
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			res = res + (string(v + 32))
		} else {
			res += string(v)
		}
	}
	return res
}
