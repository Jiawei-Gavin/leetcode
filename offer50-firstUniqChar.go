package leetcode

func firstUniqChar(s string) byte {
	hashTable := make(map[byte]int)
	for _, val := range s {
		hashTable[byte(val)]++
	}
	for _, val := range s {
		if hashTable[byte(val)] == 1 {
			return byte(val)
		}
	}
	return ' '
}
