package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}
	if root.Val <= p.Val {
		return inorderSuccessor(root.Right, p)
	} else {
		if root.Left == nil {
			return root
		}
		res := inorderSuccessor(root.Left, p)
		if res != nil {
			return res
		} else {
			return root
		}
	}
}
