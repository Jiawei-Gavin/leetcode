package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	temp1, temp2, carry, cur := 0, 0, 0, dummy
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			temp1 = 0
		} else {
			temp1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			temp2 = 0
		} else {
			temp2 = l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: (temp1 + temp2 + carry) % 10}
		cur = cur.Next
		carry = (temp1 + temp2 + carry) / 10
	}
	return dummy.Next
}
