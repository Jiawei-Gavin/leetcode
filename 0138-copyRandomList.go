package leetcode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	hashMap := make(map[*Node]*Node)
	for cur := head; cur != nil; cur = cur.Next {
		hashMap[cur] = &Node{Val: cur.Val}
	}
	for cur := head; cur != nil; cur = cur.Next {
		hashMap[cur].Next = hashMap[cur.Next]
		hashMap[cur].Random = hashMap[cur.Random]
	}
	return hashMap[head]
}
