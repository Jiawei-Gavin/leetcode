package leetcode

import (
	"sort"
	"strings"
)

// solution1
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	hashMap := make(map[rune]int)
	for _, s := range s1 {
		hashMap[s]++
	}
	for _, s := range s2 {
		hashMap[s]--
	}
	for _, v := range hashMap {
		if v > 0 {
			return false
		}
	}
	return true
}

// solution2
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	tmp1 := strings.Split(s1, "")
	tmp2 := strings.Split(s2, "")
	sort.Strings(tmp1)
	res1 := strings.Join(tmp1, "")
	sort.Strings(tmp2)
	res2 := strings.Join(tmp2, "")
	return res1 == res2
}

// solution3
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var res1 int32
	var res2 int32
	for i := range s1 {
		res1 += 1 << (s1[i] - 'a')
		res2 += 1 << (s2[i] - 'a')
	}
	return res1 == res2
}
