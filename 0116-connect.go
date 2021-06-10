package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// solution1
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		for node := leftmost; node != nil; node = node.Next {
			node.Left.Next = node.Right

			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}

// solution2
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	left := root.Left
	right := root.Right
	for left != nil {
		left.Next = right
		left = left.Right
		right = right.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}
