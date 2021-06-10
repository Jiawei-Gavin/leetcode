package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// solution1
func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var N []int
	for len(queue) != 0 {
		var tmp []*TreeNode
		tmp = append(tmp, queue...)
		queue = []*TreeNode{}
		for _, node := range tmp {
			if node.Left != nil {
				queue = append(queue, node.Left)
				if node.Left.Val == x || node.Left.Val == y {
					N = append(N, node.Val)
				}
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				if node.Right.Val == x || node.Right.Val == y {
					N = append(N, node.Val)
				}
			}
		}
		if len(N) == 1 {
			return false
		} else if len(N) == 2 {
			return N[0] != N[1]
		}
	}
	return false
}

// solution2
func isCousins(root *TreeNode, x int, y int) bool {
	var xDepth, yDepth, xParent, yParent int
	dfs(root, x, 0, -1, &xParent, &xDepth)
	dfs(root, y, 0, -1, &yParent, &yDepth)
	return xDepth > 1 && xDepth == yDepth && xParent != yParent
}

func dfs(root *TreeNode, val, depth, last int, parent, result *int) {
	if root == nil {
		return
	}
	if root.Val == val {
		*result = depth
		*parent = last
		return
	}
	depth++
	dfs(root.Left, val, depth, root.Val, parent, result)
	dfs(root.Right, val, depth, root.Val, parent, result)
}
