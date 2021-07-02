package leetcode

// solution1
func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)-1
	for left <= right {
		mid := (left + right) >> 1
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] > target {
			right = mid - 1
		} else if matrix[mid][0] < target {
			left = mid + 1
		}
	}
	if right < 0 {
		return false
	}
	row := right
	left, right = 0, len(matrix[row])-1
	for left <= right {
		mid := (left + right) >> 1
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] > target {
			right = mid - 1
		} else if matrix[row][mid] < target {
			left = mid + 1
		}
	}
	return false
}

// solution2
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n
	for left < right {
		mid := (left + right) >> 1
		if matrix[mid/n][mid%n] < target {
			left = mid + 1
		} else if matrix[mid/n][mid%n] > target {
			right = mid
		} else {
			return true
		}
	}
	return false
}
