package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	hashMap := make(map[int]int)
	for _, value := range nums {
		hashMap[value]++
	}
	var sortNums []int
	for key := range hashMap {
		sortNums = append(sortNums, key)
	}
	sort.Ints(sortNums)
	var result [][]int
	for i := 0; i < len(sortNums); i++ {
		if sortNums[i] == 0 && hashMap[sortNums[i]] >= 3 {
			result = append(result, []int{sortNums[i], sortNums[i], sortNums[i]})
		}
		for j := i + 1; j < len(sortNums); j++ {
			if sortNums[i]*2+sortNums[j] == 0 && hashMap[sortNums[i]] >= 2 {
				result = append(result, []int{sortNums[i], sortNums[i], sortNums[j]})
			}
			if sortNums[i]+sortNums[j]*2 == 0 && hashMap[sortNums[j]] >= 2 {
				result = append(result, []int{sortNums[i], sortNums[j], sortNums[j]})
			}
			remainder := 0 - sortNums[i] - sortNums[j]
			if sortNums[j] < remainder && hashMap[remainder] > 0 {
				result = append(result, []int{sortNums[i], sortNums[j], remainder})
			}
		}
	}
	return result
}
