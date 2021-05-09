package leetcode

import "math"

func increasingTriplet(nums []int) bool {
	max1, max2 := math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		if max1 >= v {
			max1 = v
		} else if max2 >= v {
			max2 = v
		} else {
			return true
		}
	}
	return false
}
