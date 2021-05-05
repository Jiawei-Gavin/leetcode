package leetcode

func isAnagram(s string, t string) bool {
	charS := [26]int{}
	for _, value := range s {
		charS[value-'a']++
	}
	charT := [26]int{}
	for _, value := range t {
		charT[value-'a']++
	}
	return charS == charT
}
