package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return nil
	}
	var res []*ListNode
	t := []*TreeNode{tree}
	for len(t) > 0 {
		var tmp []*TreeNode
		node := &ListNode{Val: 0}
		head := node
		for i := 0; i < len(t); i++ {
			head.Next = &ListNode{Val: t[i].Val}
			head = head.Next
			if t[i].Left != nil {
				tmp = append(tmp, t[i].Left)
			}
			if t[i].Right != nil {
				tmp = append(tmp, t[i].Right)
			}
		}
		t = tmp
		res = append(res, node.Next)
	}
	return res
}
