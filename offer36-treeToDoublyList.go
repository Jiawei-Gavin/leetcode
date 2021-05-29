package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre *TreeNode
var head *TreeNode

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	pre, head = nil, nil
	dfs(root)
	pre.Right, head.Left = head, pre
	return head
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if pre != nil {
		pre.Right = root
	} else {
		head = root
	}
	root.Left = pre
	pre = root
	dfs(root.Right)
}
