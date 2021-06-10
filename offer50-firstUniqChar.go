package leetcode

func firstUniqChar(s string) byte {
	hashMap := make(map[byte]int)
	for _, val := range s {
		hashMap[byte(val)]++
	}
	for _, val := range s {
		if hashMap[byte(val)] == 1 {
			return byte(val)
		}
	}
	return ' '
}
