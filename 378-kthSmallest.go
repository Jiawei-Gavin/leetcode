package leetcode

func kthSmallest(matrix [][]int, k int) int {
	left := matrix[0][0]
	right := matrix[len(matrix)-1][len(matrix[0])-1]
	for left < right {
		mid := (left + right) >> 1
		if search(matrix, mid) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func search(tmp [][]int, key int) int {
	res := 0
	for i := len(tmp) - 1; i >= 0; i-- {
		for j := 0; j < len(tmp[0]); j++ {
			if tmp[i][j] <= key {
				res++
			} else {
				break
			}
		}
	}
	return res
}
