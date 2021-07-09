package leetcode

func isIsomorphic(s string, t string) bool {
	temp1 := make(map[uint8]uint8)
	temp2 := make(map[uint8]uint8)
	for i := 0; i < len(s); i++ {
		x, ok1 := temp1[s[i]]
		y, ok2 := temp2[t[i]]
		if (ok1 && x != t[i]) || (ok2 && y != s[i]) {
			return false
		}
		temp1[s[i]] = t[i]
		temp2[t[i]] = s[i]
	}
	return true
}
