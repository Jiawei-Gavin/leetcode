package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	count := 0
	helper(root, 0, sum, &count)
	left := pathSum(root.Left, sum)
	right := pathSum(root.Right, sum)
	return count + left + right
}

func helper(root *TreeNode, cur, sum int, count *int) {
	if root == nil {
		return
	}
	cur = cur + root.Val
	if cur == sum {
		*count++
	}
	helper(root.Left, cur, sum, count)
	helper(root.Right, cur, sum, count)
}
