package leetcode

func reverseWords(s []byte) {
	reverse(s, 0, len(s)-1)
	last := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			reverse(s, last, i-1)
			last = i + 1
		}
	}
	reverse(s, last, len(s)-1)
}

func reverse(s []byte, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
