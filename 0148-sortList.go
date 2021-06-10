package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	r := sortList(slow.Next)
	slow.Next = nil
	l := sortList(head)
	tmp := &ListNode{Val: 0}
	return mergeList(tmp, l, r)
}

func mergeList(tmp, l, r *ListNode) *ListNode {
	p := tmp
	for l != nil && r != nil {
		if l.Val < r.Val {
			p.Next = l
			l = l.Next
		} else {
			p.Next = r
			r = r.Next
		}
		p = p.Next
	}
	if l == nil {
		p.Next = r
	}
	if r == nil {
		p.Next = l
	}
	return tmp.Next
}
