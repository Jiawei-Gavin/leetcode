package leetcode

import (
	"math"
	"strconv"
)

func smallestGoodBase(n string) string {
	num, _ := strconv.Atoi(n)
	for m := int(math.Log2(float64(num))); m >= 1; m-- {
		left, right := 2, int(math.Pow(float64(num), 1/float64(m)))+1
		for left < right {
			mid := (left + right) >> 1
			sum := 1
			for j := 0; j < m; j++ {
				sum = sum*mid + 1
			}
			if sum == num {
				return strconv.Itoa(mid)
			} else if sum < num {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return strconv.Itoa(num - 1)
}
