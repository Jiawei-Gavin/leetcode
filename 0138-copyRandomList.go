package leetcode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	hashMap := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		hashMap[cur] = &Node{
			Val: cur.Val,
		}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		hashMap[cur].Next = hashMap[cur.Next]
		hashMap[cur].Random = hashMap[cur.Random]
		cur = cur.Next
	}
	return hashMap[head]
}
