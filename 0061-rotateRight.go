package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		n++
	}
	if n-k%n == n {
		return head
	}
	cur.Next = head
	for i := n - k%n; i > 0; i-- {
		cur = cur.Next
	}
	res := cur.Next
	cur.Next = nil
	return res
}
