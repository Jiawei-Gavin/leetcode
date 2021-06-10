package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leaf1, leaf2 []int
	dfs(root1, &leaf1)
	dfs(root2, &leaf2)
	if len(leaf1) != len(leaf2) {
		return false
	}
	for i := range leaf1 {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, leaf *[]int) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			*leaf = append(*leaf, root.Val)
		}
		dfs(root.Left, leaf)
		dfs(root.Right, leaf)
	}
}
