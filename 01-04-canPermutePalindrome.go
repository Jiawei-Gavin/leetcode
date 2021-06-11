package leetcode

func canPermutePalindrome(s string) bool {
	hashMap := make(map[rune]int)
	for _, v := range s {
		hashMap[v] = hashMap[v] + 1
	}
	odd := 0
	for _, v := range hashMap {
		if v%2 != 0 {
			odd++
		}
	}
	return odd == 0 || odd == 1
}
