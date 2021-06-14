package leetcode

func hanota(A []int, B []int, C []int) []int {
	if A == nil {
		return nil
	}
	move(len(A), &A, &B, &C)
	return C
}

func move(n int, A *[]int, B *[]int, C *[]int) {
	if n == 0 {
		return
	}
	if n == 1 {
		*C = append(*C, (*A)[len(*A)-1])
		*A = (*A)[:len(*A)-1]
	} else {
		move(n-1, A, C, B)
		move(1, A, B, C)
		move(n-1, B, A, C)
	}
}
