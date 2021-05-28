package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result [][]int

// solution1
func levelOrder(root *TreeNode) [][]int {
	result = make([][]int, 0)
	dfs(root, 0)
	return result
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if len(result) == level {
		result = append(result, []int{root.Val})
	} else {
		result[level] = append(result[level], root.Val)
	}
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}

// solution2
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var level int
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		length := len(queue)
		result = append(result, []int{})
		for 0 < length {
			length--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			result[level] = append(result[level], queue[0].Val)
			queue = queue[1:]
		}
		level++
	}
	return result
}
