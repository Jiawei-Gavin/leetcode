package leetcode

import (
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
	r    *rand.Rand
}

func Constructor(head *ListNode) Solution {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return Solution{
		head: head,
		r:    r,
	}
}

func (this *Solution) GetRandom() int {
	i := 2
	cur, result := this.head.Next, this.head.Val
	for cur != nil {
		if this.r.Intn(i) == 0 {
			result = cur.Val
		}
		i++
		cur = cur.Next
	}
	return result
}
