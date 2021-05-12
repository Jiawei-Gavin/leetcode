package leetcode

func canJump(nums []int) bool {
	maxJump := 0
	for i := 0; i < len(nums); i++ {
		if i <= maxJump {
			maxJump = max(maxJump, i+nums[i])
			if maxJump >= len(nums)-1 {
				return true
			}
		}

	}
	return false
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
