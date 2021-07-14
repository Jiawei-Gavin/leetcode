package leetcode

func checkValidString(s string) bool {
	left, right := 0, 0
	for _, v := range s {
		if v == '(' {
			left++
			right++
		} else if v == ')' {
			if left > 0 {
				left--
			}
			right--
		} else {
			if left > 0 {
				left--
			}
			right++
		}
		if right < 0 {
			return false
		}
	}
	return left == 0
}
