package leetcode

func circularArrayLoop(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	for i := range nums {
		if nums[i] == 0 {
			continue
		}
		slow, fast := i, next(i, nums)
		for nums[slow]*nums[fast] > 0 && nums[fast]*nums[next(fast, nums)] > 0 {
			if slow == next(slow, nums) {
				break
			}
			if slow == fast {
				return true
			}
			slow = next(slow, nums)
			fast = next(fast, nums)
			fast = next(fast, nums)
		}
	}
	return false
}

func next(i int, nums []int) int {
	res := i + nums[i]
	if res >= 0 {
		return res % len(nums)
	}
	return len(nums) + res%len(nums)
}
