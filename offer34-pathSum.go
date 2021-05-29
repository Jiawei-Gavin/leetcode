package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	var result [][]int
	var path []int
	dfs(&result, root, path, target)
	return result
}

func dfs(result *[][]int, root *TreeNode, path []int, target int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && root.Val == target {
		ans := make([]int, len(path)+1)
		copy(ans, append(path, root.Val))
		*result = append(*result, ans)
		return
	}
	path = append(path, root.Val)
	dfs(result, root.Left, path, target-root.Val)
	dfs(result, root.Right, path, target-root.Val)
}
