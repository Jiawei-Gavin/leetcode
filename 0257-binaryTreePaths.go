package leetcode

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	dfs(root, "", &res)
	return res
}

func dfs(root *TreeNode, path string, res *[]string) {
	if len(path) > 0 {
		path += "->" + strconv.Itoa(root.Val)
	} else {
		path += strconv.Itoa(root.Val)
	}
	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
		return
	}
	if root.Left != nil {
		dfs(root.Left, path, res)
	}
	if root.Right != nil {
		dfs(root.Right, path, res)
	}
}
