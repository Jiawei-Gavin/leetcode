package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		left, right := 0, len(matrix[0])-1
		for left <= right {
			mid := left + (right-left)>>1
			if row[mid] > target {
				right = mid - 1
			} else if row[mid] < target {
				left = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}
