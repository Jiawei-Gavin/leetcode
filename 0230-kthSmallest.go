package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// solution1 inorder
func kthSmallest(root *TreeNode, k int) int {
	var arr []int
	res := inorder(root, &arr)
	return res[k-1]
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

// solution2
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			k--
			if k == 0 {
				return root.Val
			}
			root = root.Right
		}
	}
	return 0
}
