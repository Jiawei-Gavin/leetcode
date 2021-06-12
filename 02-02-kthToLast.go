package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	fast, slow := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	return slow.Val
}
