package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		} else {
			left := dfs(node.Left)
			right := dfs(node.Right)
			res = max(left+right, res)
			return max(left, right) + 1
		}
	}
	dfs(root)
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
