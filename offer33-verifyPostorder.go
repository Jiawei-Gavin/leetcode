package leetcode

func verifyPostorder(postorder []int) bool {
	return recur(postorder, 0, len(postorder)-1)
}

func recur(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}
	index := i
	for postorder[index] < postorder[j] {
		index++
	}
	mid := index
	for postorder[index] > postorder[j] {
		index++
	}
	return index == j && recur(postorder, i, mid-1) && recur(postorder, mid, j-1)
}
