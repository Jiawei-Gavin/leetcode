package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	tail := head
	for i := 0; i < k; i++ {
		if tail == nil {
			return head
		}
		tail = tail.Next
	}
	cur := reverse(head, tail)
	head.Next = reverseKGroup(tail, k)
	return cur
}

func reverse(head, tail *ListNode) *ListNode {
	pre := &ListNode{}
	for head != tail {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
