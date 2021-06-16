package leetcode

import "sort"

func maxAliveYear(birth []int, death []int) int {
	sort.Ints(birth)
	sort.Ints(death)
	year := birth[0]
	end, max, num := 0, 0, 0
	for i := 0; i < len(birth); i++ {
		num++
		for birth[i] > death[end] {
			num--
			end++
		}
		if num > max {
			year = birth[i]
			max = num
		}
	}
	return year
}
