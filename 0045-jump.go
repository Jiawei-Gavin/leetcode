package leetcode

// solution1
func jump(nums []int) int {
	var res int
	index := len(nums) - 1
	for index > 0 {
		for i := 0; i < index; i++ {
			if i+nums[i] >= index {
				index = i
				res++
				break
			}
		}
	}
	return res
}

// solution2
func jump(nums []int) int {
	end, maxPosition, res := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, nums[i]+i)
		if i == end {
			end = maxPosition
			res++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
