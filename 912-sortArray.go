package leetcode

// solution1
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left < right {
		nums[right], nums[(left+right)>>1] = nums[(left+right)>>1], nums[right]
		i := left - 1
		for j := left; j < right; j++ {
			if nums[j] < nums[right] {
				i++
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		i++
		nums[i], nums[right] = nums[right], nums[i]
		quickSort(nums, left, i-1)
		quickSort(nums, i+1, right)
	}
}

// solution2
func sortArray(nums []int) []int {
	MergeSort(nums, 0, len(nums)-1)
	return nums
}

func MergeSort(nums []int, left, right int) {
	if left < right {
		mid := (left + right) >> 1
		MergeSort(nums, left, mid)
		MergeSort(nums, mid+1, right)
		i, j := left, mid+1
		var tmp []int
		for i <= mid || j <= right {
			if i > mid || (j <= right && nums[j] < nums[i]) {
				tmp = append(tmp, nums[j])
				j++
			} else {
				tmp = append(tmp, nums[i])
				i++
			}
		}
		copy(nums[left:right+1], tmp)
	}
}
