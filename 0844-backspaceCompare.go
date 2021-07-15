package leetcode

func backspaceCompare(s string, t string) bool {
	p1, p2 := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				p1++
				i--
			} else if p1 > 0 {
				p1--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				p2++
				j--
			} else if p2 > 0 {
				p2--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else {
			if i >= 0 || j >= 0 {
				return false
			}
		}
		i--
		j--
	}
	return true
}
