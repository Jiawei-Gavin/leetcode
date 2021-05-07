package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid+1:]
	return &TreeNode{nums[mid], sortedArrayToBST(left), sortedArrayToBST(right)}
}
