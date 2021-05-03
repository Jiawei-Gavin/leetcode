package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	hashTable := make(map[int]int)
	for _, value := range nums {
		hashTable[value]++
	}
	var sortNums []int
	for key := range hashTable {
		sortNums = append(sortNums, key)
	}
	sort.Ints(sortNums)
	var result [][]int
	for i := 0; i < len(sortNums); i++ {
		if sortNums[i] == 0 && hashTable[sortNums[i]] >= 3 {
			result = append(result, []int{sortNums[i], sortNums[i], sortNums[i]})
		}
		for j := i + 1; j < len(sortNums); j++ {
			if sortNums[i]*2+sortNums[j] == 0 && hashTable[sortNums[i]] >= 2 {
				result = append(result, []int{sortNums[i], sortNums[i], sortNums[j]})
			}
			if sortNums[i]+sortNums[j]*2 == 0 && hashTable[sortNums[j]] >= 2 {
				result = append(result, []int{sortNums[i], sortNums[j], sortNums[j]})
			}
			remainder := 0 - sortNums[i] - sortNums[j]
			if sortNums[j] < remainder && hashTable[remainder] > 0 {
				result = append(result, []int{sortNums[i], sortNums[j], remainder})
			}
		}
	}
	return result
}
