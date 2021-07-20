package leetcode

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		temp1, temp2 := nums[left]*nums[left], nums[right]*nums[right]
		if temp1 > temp2 {
			res[i] = temp1
			left++
		} else {
			res[i] = temp2
			right--
		}
	}
	return res
}
