package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// solution1
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

// solution2
func rangeSumBST(root *TreeNode, low int, high int) int {
	queue, sum := []*TreeNode{root}, 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		if node.Val > high {
			queue = append(queue, node.Left)
		} else if node.Val < low {
			queue = append(queue, node.Right)
		} else {
			sum += node.Val
			queue = append(queue, node.Left, node.Right)
		}
	}
	return sum
}
