package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func plusOne(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	cur := dummy
	for head != nil {
		if head.Val != 9 {
			cur = head
		}
		head = head.Next
	}
	cur.Val++
	cur = cur.Next
	for cur != nil {
		cur.Val = 0
		cur = cur.Next
	}

	if dummy.Val != 0 {
		return dummy
	} else {
		return dummy.Next
	}
}
