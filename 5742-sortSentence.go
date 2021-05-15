package leetcode

import (
	"sort"
	"strings"
)

// solution1
func sortSentence(s string) string {
	res := ""
	hashTable := make(map[int]string)
	array := strings.Fields(s)
	for i := 0; i < len(array); i++ {
		key := int(array[i][len(array[i])-1]) - 48
		hashTable[key] = array[i][:len(array[i])-1]
	}
	for i := 1; i <= len(array); i++ {
		res += hashTable[i]
		if i != len(array) {
			res += " "
		}
	}
	return res
}

// solution2
func sortSentence(s string) (ans string) {
	a := strings.Split(s, " ")
	sort.Slice(a, func(i, j int) bool { return a[i][len(a[i])-1] < a[j][len(a[j])-1] })
	for i, v := range a {
		a[i] = v[:len(s)-1]
	}
	return strings.Join(a, " ")
}
