package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var preSlow *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	root := &TreeNode{Val: slow.Val}
	if preSlow != nil {
		preSlow.Next = nil
		root.Left = sortedListToBST(head)
	}
	root.Right = sortedListToBST(slow.Next)
	return root
}
