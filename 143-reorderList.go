package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// solution1
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	var list []*ListNode
	p := head
	for p != nil {
		list = append(list, p)
		p = p.Next
	}

	start, end := 0, len(list)-1
	for start < end {
		list[start].Next = list[end]
		list[end].Next = list[start+1]
		start++
		end--
	}
	list[start].Next = nil
}

// solution2
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow
	p := mid.Next
	for p.Next != nil {
		current := p.Next
		p.Next = current.Next
		current.Next = mid.Next
		mid.Next = current
	}

	p1 := head
	p2 := slow.Next
	for p1 != mid {
		mid.Next = p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p2.Next
		p2 = mid.Next
	}
	return
}
