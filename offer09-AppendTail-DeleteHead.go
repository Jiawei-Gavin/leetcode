package leetcode

type CQueue struct {
	stack1, stack2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stack2) == 0 {
		if len(this.stack1) == 0 {
			return -1
		}
		for len(this.stack1) > 0 {
			index := len(this.stack1) - 1
			value := this.stack1[index]
			this.stack2 = append(this.stack2, value)
			this.stack1 = this.stack1[:index]
		}
	}
	index := len(this.stack2) - 1
	value := this.stack2[index]
	this.stack2 = this.stack2[:index]
	return value
}
