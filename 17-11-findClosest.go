package leetcode

import "math"

func findClosest(words []string, word1 string, word2 string) int {
	res := math.MaxInt64
	index1, index2 := -1, -1
	for i, word := range words {
		if word == word1 {
			index1 = i
		} else if word == word2 {
			index2 = i
		}
		if index1 != -1 && index2 != -1 {
			diff := abs(index1 - index2)
			if diff < res {
				res = diff
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
