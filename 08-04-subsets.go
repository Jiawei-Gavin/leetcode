package leetcode

func subsets(nums []int) [][]int {
	var res [][]int
	var dfs func(path []int, index int)
	dfs = func(path []int, index int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(path, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs([]int{}, 0)
	return res
}
