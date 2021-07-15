package leetcode

func reverseOnlyLetters(s string) string {
	var res string
	j := len(s) - 1
	for _, num := range s {
		if isLetter(num) {
			for !isLetter(rune(s[j])) {
				j--
			}
			res += string(s[j])
			j--
		} else {
			res += string(num)
		}
	}
	return res
}

func isLetter(s rune) bool {
	if (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z') {
		return true
	}
	return false
}
