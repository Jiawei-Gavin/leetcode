package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}

func dfs(root *TreeNode) int {
	res := 0
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			res += root.Left.Val
		} else {
			res += dfs(root.Left)
		}
	}
	if root.Right != nil {
		res += dfs(root.Right)
	}
	return res
}
