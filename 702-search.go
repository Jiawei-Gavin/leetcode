package leetcode

func search(reader ArrayReader, target int) int {
	left, right := 0, 1000001
	for left <= right {
		mid := left + (right-left)>>1
		if reader.get(mid) == target {
			return mid
		}

		if reader.get(mid) > target {
			right = mid - 1
			continue
		}
		if reader.get(mid) < target {
			left = mid + 1
			continue
		}
	}
	return -1
}
