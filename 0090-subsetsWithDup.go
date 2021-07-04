package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	dfs(nums, 0, []int{}, &res)
	return res
}

func dfs(nums []int, index int, subset []int, res *[][]int) {
	temp := make([]int, len(subset))
	copy(temp, subset)
	*res = append(*res, temp)
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		subset = append(subset, nums[i])
		dfs(nums, i+1, subset, res)
		subset = subset[:len(subset)-1]
	}
}
