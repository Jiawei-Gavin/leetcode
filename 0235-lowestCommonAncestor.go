package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	result := root
	for {
		if p.Val < result.Val && q.Val < result.Val {
			result = result.Left
		} else if p.Val > result.Val && q.Val > result.Val {
			result = result.Right
		} else {
			return result
		}
	}
}
