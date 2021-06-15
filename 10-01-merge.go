package leetcode

func merge(A []int, m int, B []int, n int) {
	for i := m + n - 1; i >= 0; i-- {
		if n-1 < 0 {
			return
		}
		if m-1 < 0 {
			n--
			A[i] = B[n]
			continue
		}
		if A[m-1] < B[n-1] {
			n--
			A[i] = B[n]
		} else {
			m--
			A[i] = A[m]
		}
	}
}
