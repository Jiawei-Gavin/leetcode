package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	pre, cur := head, head
	for i := 0; i < k; i++ {
		if pre == nil {
			return nil
		}
		pre = pre.Next
	}
	for pre != nil {
		pre, cur = pre.Next, cur.Next
	}
	return cur
}
