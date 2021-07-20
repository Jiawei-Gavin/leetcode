package leetcode

func lengthOfLongestSubstring(s string) int {
	hashMap := make(map[byte]int)
	res := 0
	for start, end := 0, 0; end < len(s); end++ {
		if i, ok := hashMap[s[end]]; ok {
			start = max(i, start)
		}
		hashMap[s[end]] = end + 1
		res = max(res, end-start+1)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
