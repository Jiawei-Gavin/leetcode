package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	var res int
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum = sum<<1 + root.Val
		if root.Left == nil && root.Right == nil {
			res += sum
			return
		}
		dfs(root.Left, sum)
		dfs(root.Right, sum)
	}
	dfs(root, 0)
	return res
}
