package leetcode

type MyCircularQueue struct {
	q                        []int
	maxLen, front, rear, len int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		q:      make([]int, k),
		maxLen: k,
		front:  -1,
		rear:   0,
		len:    0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.len >= this.maxLen {
		return false
	}
	this.front = (this.front + 1) % this.maxLen
	this.q[this.front] = value
	this.len++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.len == 0 {
		return false
	}
	this.rear = (this.rear + 1) % this.maxLen
	this.len--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.len == 0 {
		return -1
	}
	return this.q[this.rear]
}

func (this *MyCircularQueue) Rear() int {
	if this.len == 0 {
		return -1
	}
	return this.q[this.front]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.len == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.maxLen == this.len
}
