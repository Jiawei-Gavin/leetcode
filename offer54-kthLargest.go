package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var result int
	dfs(root, &result, &k)
	return result
}

func dfs(root *TreeNode, result *int, k *int) {
	if root == nil {
		return
	}
	dfs(root.Right, result, k)
	if *k == 0 {
		return
	}
	*k--
	if *k == 0 {
		*result = root.Val
	}
	dfs(root.Left, result, k)
}
