package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var combine []int
	var res [][]int
	sort.Ints(candidates)
	var dfs func(target, index int)
	dfs = func(target, index int) {
		if target == 0 {
			res = append(res, append([]int(nil), combine...))
			return
		}
		for i := index; i < len(candidates); i++ {
			if target < candidates[i] {
				break
			}
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			combine = append(combine, candidates[i])
			dfs(target-candidates[i], i+1)
			combine = combine[:len(combine)-1]
		}
	}
	dfs(target, 0)
	return res
}
