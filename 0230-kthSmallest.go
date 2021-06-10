package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var arr []int
	result := inorder(root, &arr)
	return result[k-1]
}

func inorder(root *TreeNode, arr *[]int) []int {
	if root == nil {
		return *arr
	}
	inorder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorder(root.Right, arr)
	return *arr
}
