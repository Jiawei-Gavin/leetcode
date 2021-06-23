package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// solution1
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, right, left int) bool {
	if root == nil {
		return true
	}
	if root.Val <= right || root.Val >= left {
		return false
	}
	return helper(root.Left, right, root.Val) && helper(root.Right, root.Val, left)
}

// solution2
func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}
