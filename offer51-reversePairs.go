package leetcode

func reversePairs(nums []int) int {
	result := mergeSort(nums, 0, len(nums)-1)
	return result
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := (start + end) / 2
	count := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	var tmp []int
	left, right := start, mid+1
	for left <= mid && right <= end {
		if nums[left] <= nums[right] {
			tmp = append(tmp, nums[left])
			count += right - (mid + 1)
			left++
		} else {
			tmp = append(tmp, nums[right])
			right++
		}
	}
	for ; left <= mid; left++ {
		tmp = append(tmp, nums[left])
		count += end - (mid + 1) + 1
	}
	for ; right <= end; right++ {
		tmp = append(tmp, nums[right])
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}
	return count
}
