package leetcode

func repeatedSubstringPattern(s string) bool {
	for i := 1; i*2 <= len(s); i++ {
		if len(s)%i == 0 {
			match := true
			for j := i; j < len(s); j++ {
				if s[j] != s[j-i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}
