package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	largeDummy, smallDummy := &ListNode{0, nil}, &ListNode{0, nil}
	large, small := largeDummy, smallDummy
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = largeDummy.Next
	return smallDummy.Next
}
