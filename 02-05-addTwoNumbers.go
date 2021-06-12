package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	result := &ListNode{}
	dummy := result
	tmp := 0
	for l1 != nil || l2 != nil {
		res := tmp
		if l1 != nil {
			res += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			res += l2.Val
			l2 = l2.Next
		}
		result.Next = &ListNode{Val: res % 10}
		tmp = res / 10
		result = result.Next
	}
	if tmp != 0 {
		result.Next = &ListNode{Val: tmp}
	}
	return dummy.Next
}
