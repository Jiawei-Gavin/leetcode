package leetcode

func oneEditAway(first string, second string) bool {
	if first == second {
		return true
	}
	l1, l2 := len(first), len(second)
	if l1 < l2 {
		return oneEditAway(second, first)
	} else if l1-l2 > 1 {
		return false
	}

	l, r1, r2 := 0, l1-1, l2-1
	for l <= r1 && l <= r2 && first[l] == second[l] {
		l++
	}
	for r1 >= 0 && r2 >= 0 && first[r1] == second[r2] {
		r1--
		r2--
	}
	return r1-l < 1 && r2-l < 1
}
