package leetcode

func isCovered(ranges [][]int, left int, right int) bool {
	for i := left; i <= right; i++ {
		res := false
		for _, v := range ranges {
			if i >= v[0] && i <= v[1] {
				res = true
				break
			}
		}
		if !res {
			return false
		}
	}
	return true
}
