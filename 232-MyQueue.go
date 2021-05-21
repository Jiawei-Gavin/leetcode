package leetcode

type MyQueue struct {
	Stack, Queue []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.Stack = append(this.Stack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.Queue) == 0 {
		this.fromStackToQueue()
	}
	popped := (this.Queue)[len(this.Queue)-1]
	this.Queue = (this.Queue)[:len(this.Queue)-1]
	return popped
}

func (this *MyQueue) fromStackToQueue() {
	for len(this.Stack) > 0 {
		popped := (this.Stack)[len(this.Stack)-1]
		this.Stack = (this.Stack)[:len(this.Stack)-1]
		this.Queue = append(this.Queue, popped)
	}
}

func (this *MyQueue) Peek() int {
	if len(this.Queue) == 0 {
		this.fromStackToQueue()
	}
	return (this.Queue)[len(this.Queue)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.Stack)+len(this.Queue) == 0
}
