package leetcode

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(rec1[2] <= rec2[0] || rec2[2] <= rec1[0]) && !(rec1[3] <= rec2[1] || rec2[3] <= rec1[1])
}
