package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{preorder[0], nil, nil}
	for i := 1; i < len(preorder); i++ {
		buildBST(preorder[i], root)
	}
	return root
}

func buildBST(val int, root *TreeNode) {
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
			return
		} else {
			buildBST(val, root.Left)
		}
	}
	if val > root.Val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
			return
		} else {
			buildBST(val, root.Right)
		}
	}
}
