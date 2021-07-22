package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	nums []int
	root *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var nums []int
	inorder(root, &nums)
	return BSTIterator{
		nums: nums,
		root: root,
	}
}

func inorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inorder(root.Right, nums)
}

func (this *BSTIterator) Next() int {
	res := this.nums[0]
	this.nums = this.nums[1:]
	return res
}

func (this *BSTIterator) HasNext() bool {
	if len(this.nums) > 0 {
		return true
	}
	return false
}
