package leetcode

func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}

func maxPalindrome(s string, left, right int, res string) string {
	var sub string
	for left >= 0 && right < len(s) && s[left] == s[right] {
		sub = s[left : right+1]
		left--
		right++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}
