package leetcode

import "sort"

func topKFrequent(words []string, k int) []string {
	hashTable := map[string]int{}
	for _, word := range words {
		hashTable[word]++
	}
	result := make([]string, 0, len(hashTable))
	for v := range hashTable {
		result = append(result, v)
	}

	sort.Slice(result, func(i, j int) bool {
		s, t := result[i], result[j]
		return hashTable[s] > hashTable[t] || hashTable[s] == hashTable[t] && s < t
	})
	return result[:k]
}
