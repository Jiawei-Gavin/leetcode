package leetcode

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

type LRUCache struct {
	cache     map[int]*Node
	size      int
	capacity  int
	dummyHead *Node
	dummyTail *Node
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		cache:     map[int]*Node{},
		size:      0,
		capacity:  capacity,
		dummyHead: &Node{0, 0, nil, nil},
		dummyTail: &Node{0, 0, nil, nil},
	}
	lruCache.dummyHead.Next = lruCache.dummyTail
	lruCache.dummyTail.Pre = lruCache.dummyHead
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := &Node{key, value, nil, nil}
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			tail := this.removeTail()
			delete(this.cache, tail.Key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.Val = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *Node) {
	node.Pre = this.dummyHead
	node.Next = this.dummyHead.Next
	this.dummyHead.Next.Pre = node
	this.dummyHead.Next = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *Node {
	node := this.dummyTail.Pre
	this.removeNode(node)
	return node
}
