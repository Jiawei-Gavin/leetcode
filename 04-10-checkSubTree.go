package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil {
		return t2 == nil
	}
	return isSame(t1, t2) || checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2)
}

func isSame(root, p *TreeNode) bool {
	if root == nil && p == nil {
		return true
	}
	if root == nil || p == nil {
		return false
	}
	return root.Val == p.Val && isSame(root.Left, p.Left) && isSame(root.Right, p.Right)
}
