package leetcode

func multiSearch(big string, smalls []string) [][]int {
	res := make([][]int, len(smalls))
	for i, small := range smalls {
		if small == "" {
			res[i] = []int{}
			continue
		}
		for j, _ := range big {
			if j+len(small) <= len(big) && big[j:j+len(small)] == small {
				res[i] = append(res[i], j)
			}
		}
	}
	return res
}
