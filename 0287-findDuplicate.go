package leetcode

// solution1
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return slow
}

// solution2
func findDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	var res int
	for left <= right {
		mid := left + (right-left)>>1
		cnt := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			left = mid + 1
		} else {
			right = mid - 1
			res = mid
		}
	}
	return res
}
