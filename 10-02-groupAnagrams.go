package leetcode

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	hashMap := map[string][]string{}
	for _, s := range strs {
		tmp := strings.Split(s, "")
		sort.Strings(tmp)
		key := strings.Join(tmp, "")
		if _, ok := hashMap[strings.Join(tmp, "")]; ok {
			hashMap[key] = append(hashMap[key], s)
		} else {
			hashMap[key] = []string{s}
		}
	}
	for _, v := range hashMap {
		result = append(result, v)
	}
	return result
}
