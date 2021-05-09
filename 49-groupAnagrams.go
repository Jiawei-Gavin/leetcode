package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	hashtable := map[string][]string{}
	var result [][]string
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		hashtable[string(s)] = append(hashtable[string(s)], str)
	}

	for _, v := range hashtable {
		result = append(result, v)
	}
	return result
}
