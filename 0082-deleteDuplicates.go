package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// solution1
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{0, head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			val := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// solution2
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val == head.Next.Val {
		for head.Next != nil && head.Next.Val == head.Val {
			head = head.Next
		}
		return deleteDuplicates(head.Next)
	} else {
		head.Next = deleteDuplicates(head.Next)
		return head
	}
}
