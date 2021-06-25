package leetcode

type Node struct {
	Val      int
	Children []*Node
}

// solution1
func preorder(root *Node) []int {
	var res []int
	var dfs func(*Node)
	dfs = func(root *Node) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		for _, children := range root.Children {
			dfs(children)
		}
	}
	dfs(root)
	return res
}

// solution2
func preorder(root *Node) []int {
	var res []int
	var stack = []*Node{root}
	if root == nil {
		return res
	}
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		for i := len(root.Children) - 1; i >= 0; i-- {
			stack = append(stack, root.Children[i])
		}
	}
	return res
}
