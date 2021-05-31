package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// solution1
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{0, head}
	pre := dummy
	for head != nil {
		next := head.Next
		if head.Val == val {
			pre.Next = next
			break
		}
		pre = head
		head = next
	}
	return dummy.Next
}

// solution2
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}
	pre, cur := head, head.Next
	for cur != nil && cur.Val != val {
		pre, cur = cur, cur.Next
	}
	if cur != nil {
		pre.Next = cur.Next
	}
	return head
}
