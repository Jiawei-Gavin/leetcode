package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

var pre *Node
var head *Node

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}
	pre, head = nil, nil
	dfs(root)
	pre.Right, head.Left = head, pre
	return head
}

func dfs(root *Node) {
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
