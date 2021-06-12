package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow := head
	for slow != nil {
		fast := slow
		for fast.Next != nil {
			if slow.Val == fast.Next.Val {
				fast.Next = fast.Next.Next
			} else {
				fast = fast.Next
			}
		}
		slow = slow.Next
	}
	return head
}
