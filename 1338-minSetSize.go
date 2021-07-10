package leetcode

import "sort"

func minSetSize(arr []int) int {
	hashMap := make(map[int]int)
	for _, v := range arr {
		hashMap[v]++
	}
	res, cur, index, num := 0, 0, 0, len(arr)
	ans := make([]int, len(hashMap))
	for _, v := range hashMap {
		ans[index] = v
		index++
	}
	sort.Ints(ans)
	for i := len(ans) - 1; i >= 0; i-- {
		cur += ans[i]
		num -= ans[i]
		res++
		if cur >= num {
			return res
		}
	}
	return res
}
