package leetcode

func longestPalindrome(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		result = maxPalindrome(s, i, i, result)
		result = maxPalindrome(s, i, i+1, result)
	}
	return result
}

func maxPalindrome(s string, i, j int, result string) string {
	var sub string
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(result) < len(sub) {
		return sub
	}
	return result
}
