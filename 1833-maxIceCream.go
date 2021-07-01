package leetcode

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	var res int
	for _, cost := range costs {
		coins -= cost
		if coins < 0 {
			return res
		}
		res++
	}
	return res
}
