package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		length := len(queue)
		for 0 < length {
			length--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			result = append(result, queue[0].Val)
			queue = queue[1:]
		}
	}
	return result
}
