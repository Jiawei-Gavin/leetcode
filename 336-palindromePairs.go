package leetcode

func palindromePairs(words []string) [][]int {
	var result [][]int
	hashTable := make(map[string]int)
	for i := 0; i < len(words); i++ {
		hashTable[words[i]] = i
	}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i])+1; j++ {
			left, right := words[i][:j], words[i][j:]
			if v, ok := hashTable[reverse(left)]; ok && v != i && isPalindrome(right) {
				result = append(result, []int{i, v})
			}
			if v, ok := hashTable[reverse(right)]; ok && v != i && j > 0 && isPalindrome(left) {
				result = append(result, []int{v, i})
			}
		}
	}
	return result
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func reverse(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return string(b)
}
