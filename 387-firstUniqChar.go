package leetcode

func firstUniqChar(s string) int {
	char := [26]int{}
	for _, value := range s {
		char[value-'a']++
	}
	for key, value := range s {
		if char[value-'a'] == 1 {
			return key
		}
	}
	return -1
}
