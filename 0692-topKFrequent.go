package leetcode

import "sort"

func topKFrequent(words []string, k int) []string {
	hashMap := map[string]int{}
	for _, word := range words {
		hashMap[word]++
	}
	result := make([]string, 0, len(hashMap))
	for v := range hashMap {
		result = append(result, v)
	}

	sort.Slice(result, func(i, j int) bool {
		s, t := result[i], result[j]
		return hashMap[s] > hashMap[t] || hashMap[s] == hashMap[t] && s < t
	})
	return result[:k]
}
