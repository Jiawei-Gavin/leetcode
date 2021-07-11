package leetcode

func hIndex(citations []int) int {
	n := len(citations)
	left, right := 0, len(citations)-1
	for left <= right {
		mid := (left + right) >> 1
		if citations[mid] == n-mid {
			return n - mid
		} else if citations[mid] < n-mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return n - left
}
