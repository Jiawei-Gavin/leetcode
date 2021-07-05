package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// solution1
func rightSideView(root *TreeNode) []int {
	var res []int
	dfs(root, 1, &res)
	return res
}

func dfs(root *TreeNode, level int, res *[]int) {
	if root == nil {
		return
	}
	if level > len(*res) {
		*res = append(*res, root.Val)
	}
	dfs(root.Right, level+1, res)
	dfs(root.Left, level+1, res)
}

// solution2
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		res = append(res, queue[length-1].Val)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}
