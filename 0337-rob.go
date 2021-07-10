package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	res := dfs(root)
	return max(res[0], res[1])
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	dp := make([]int, 2)
	dp[0] = max(left[0], left[1]) + max(right[0], right[1])
	dp[1] = root.Val + left[0] + right[0]
	return dp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
