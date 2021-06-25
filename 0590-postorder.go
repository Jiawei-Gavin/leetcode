package leetcode

type Node struct {
	Val      int
	Children []*Node
}

// solution1
func postorder(root *Node) []int {
	var res []int
	var dfs func(*Node)
	dfs = func(root *Node) {
		if root == nil {
			return
		}
		for _, children := range root.Children {
			dfs(children)
		}
		res = append(res, root.Val)
	}
	dfs(root)
	return res
}

// solution2
func postorder(root *Node) []int {
	var res []int
	var stack = []*Node{root}
	if root == nil {
		return res
	}
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append([]int{root.Val}, res...)
		for i := 0; i < len(root.Children); i++ {
			stack = append(stack, root.Children[i])
		}
	}
	return res
}
