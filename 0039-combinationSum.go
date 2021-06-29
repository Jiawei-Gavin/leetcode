package leetcode

func combinationSum(candidates []int, target int) [][]int {
	var combine []int
	var res [][]int
	var dfs func(target, index int)
	dfs = func(target, index int) {
		if index == len(candidates) {
			return
		}
		if target == 0 {
			res = append(res, append([]int(nil), combine...))
			return
		}
		dfs(target, index+1)
		if target-candidates[index] >= 0 {
			combine = append(combine, candidates[index])
			dfs(target-candidates[index], index)
			combine = combine[:len(combine)-1]
		}
	}
	dfs(target, 0)
	return res
}
