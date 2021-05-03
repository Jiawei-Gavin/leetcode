package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
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
		if sortNums[i]*4 == target && hashTable[sortNums[i]] >= 4 {
			result = append(result, []int{sortNums[i], sortNums[i], sortNums[i], sortNums[i]})
		}
		for j := i + 1; j < len(sortNums); j++ {
			if sortNums[i]*3+sortNums[j] == target && hashTable[sortNums[i]] >= 3 {
				result = append(result, []int{sortNums[i], sortNums[i], sortNums[i], sortNums[j]})
			}
			if sortNums[i]+sortNums[j]*3 == target && hashTable[sortNums[j]] >= 3 {
				result = append(result, []int{sortNums[i], sortNums[j], sortNums[j], sortNums[j]})
			}
			if sortNums[i]*2+sortNums[j]*2 == target && hashTable[sortNums[i]] >= 2 && hashTable[sortNums[j]] >= 2 {
				result = append(result, []int{sortNums[i], sortNums[i], sortNums[j], sortNums[j]})
			}
			for k := j + 1; k < len(sortNums); k++ {
				if sortNums[i]*2+sortNums[j]+sortNums[k] == target && hashTable[sortNums[i]] >= 2 {
					result = append(result, []int{sortNums[i], sortNums[i], sortNums[j], sortNums[k]})
				}
				if sortNums[i]+sortNums[j]*2+sortNums[k] == target && hashTable[sortNums[j]] >= 2 {
					result = append(result, []int{sortNums[i], sortNums[j], sortNums[j], sortNums[k]})
				}
				if sortNums[i]+sortNums[j]+sortNums[k]*2 == target && hashTable[sortNums[k]] >= 2 {
					result = append(result, []int{sortNums[i], sortNums[j], sortNums[k], sortNums[k]})
				}
				remainder := target - sortNums[i] - sortNums[j] - sortNums[k]
				if sortNums[k] < remainder && hashTable[remainder] > 0 {
					result = append(result, []int{sortNums[i], sortNums[j], sortNums[k], remainder})
				}
			}
		}
	}
	return result
}
