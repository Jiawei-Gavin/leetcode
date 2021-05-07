package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// solution1
func isValidBST(root *TreeNode) bool {
	return ValidBST(root, math.MinInt64, math.MaxInt64)
}

func ValidBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return ValidBST(root.Left, min, root.Val) && ValidBST(root.Right, root.Val, max)
}

// solution2
func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode
	min := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= min {
			return false
		}
		min = root.Val
		root = root.Right
	}
	return true
}
