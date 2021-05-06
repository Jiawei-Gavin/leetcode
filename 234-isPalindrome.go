package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// solution1
func isPalindrome(head *ListNode) bool {
	var Lists []int
	for ; head != nil; head = head.Next {
		Lists = append(Lists, head.Val)
	}
	n := len(Lists)
	for key, value := range Lists[:n/2] {
		if value != Lists[n-1-key] {
			return false
		}
	}
	return true
}

// solution2
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fastNode := head
	slowNode := head
	for fastNode.Next != nil && fastNode.Next.Next != nil {
		fastNode = fastNode.Next.Next
		slowNode = slowNode.Next
	}

	preMiddle := slowNode
	preCurrent := slowNode.Next
	for preCurrent.Next != nil {
		current := preCurrent.Next
		preCurrent.Next = current.Next
		current.Next = preMiddle.Next
		preMiddle.Next = current
	}
	right := preMiddle.Next
	left := head
	res := true
	for right != nil {
		if left.Val != right.Val {
			res = false
			break
		}
		left = left.Next
		right = right.Next
	}
	return res
}
