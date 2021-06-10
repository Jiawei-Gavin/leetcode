package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// solution1
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1Length, l2Length int
	node1, node2 := l1, l2
	for node1 != nil {
		l1Length++
		node1 = node1.Next
	}
	for node2 != nil {
		l2Length++
		node2 = node2.Next
	}
	newHead := &ListNode{1, nil}
	if l1Length < l2Length {
		newHead.Next = add(l2, l1, l2Length-l1Length)
	} else {
		newHead.Next = add(l1, l2, l1Length-l2Length)
	}
	if newHead.Next.Val > 9 {
		newHead.Next.Val = newHead.Next.Val % 10
		return newHead
	}
	return newHead.Next
}

func add(l1 *ListNode, l2 *ListNode, offset int) *ListNode {
	if l1 == nil {
		return nil
	}
	var (
		result, node *ListNode
	)
	if offset == 0 {
		result = &ListNode{Val: l1.Val + l2.Val, Next: nil}
		node = add(l1.Next, l2.Next, 0)
	} else {
		result = &ListNode{Val: l1.Val, Next: nil}
		node = add(l1.Next, l2, offset-1)
	}
	if node != nil && node.Val > 9 {
		result.Val++
		node.Val = node.Val % 10
	}
	result.Next = node
	return result
}

// solution2
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2, carry := pushStack(l1), pushStack(l2), 0
	var newHead *ListNode
	for len(*s1) > 0 || len(*s2) > 0 {
		carry += popStack(s1) + popStack(s2)
		newHead, carry = &ListNode{Val: carry % 10, Next: newHead}, carry/10
	}
	if carry != 0 {
		newHead = &ListNode{Val: carry, Next: newHead}
	}
	return newHead
}

func pushStack(head *ListNode) *[]int {
	stack := new([]int)
	for ptr := head; ptr != nil; ptr = ptr.Next {
		*stack = append(*stack, ptr.Val)
	}
	return stack
}

func popStack(stack *[]int) (pop int) {
	if len(*stack) > 0 {
		pop = (*stack)[len(*stack)-1]
		*stack = (*stack)[0 : len(*stack)-1]
	}
	return
}
