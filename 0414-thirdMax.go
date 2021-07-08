package leetcode

import "math"

func thirdMax(nums []int) int {
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if max3 >= num || max2 == num || max1 == num {
			continue
		}
		if max1 < num {
			max3 = max2
			max2 = max1
			max1 = num
		} else if max2 < num {
			max3 = max2
			max2 = num
		} else {
			max3 = num
		}
	}
	if max3 == math.MinInt64 {
		return max1
	}
	return max3
}
