package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	sort.Ints(nums)
	backtrack([]int{}, nums, used, &result)
	return result
}

func backtrack(path, nums []int, used []bool, result *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i-1 >= 0 && nums[i-1] == nums[i] && !used[i-1] {
			continue
		}
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backtrack(path, nums, used, result)
		path = path[0 : len(path)-1]
		used[i] = false
	}
}
