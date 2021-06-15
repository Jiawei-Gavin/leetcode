package leetcode

func search(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) >> 1
		if arr[left] == target {
			return left
		} else if arr[mid] == target {
			right = mid
		} else if arr[mid] > arr[left] {
			if arr[left] <= target && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if arr[mid] < arr[right] {
			if arr[mid] < target && target <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			for left <= right && arr[mid] == arr[left] {
				left++
			}
			for left <= right && arr[right] == arr[mid] {
				right--
			}
		}
	}
	return -1
}
