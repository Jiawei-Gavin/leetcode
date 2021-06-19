package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBiNode(root *TreeNode) *TreeNode {
	head := &TreeNode{}
	curr := head
	changeNode(root, curr)
	return head.Right
}

func changeNode(root, cur *TreeNode) *TreeNode {
	if root != nil {
		cur = changeNode(root.Left, cur)
		root.Left = nil
		cur.Right = root
		cur = root
		cur = changeNode(root.Right, cur)
	}
	return cur
}
