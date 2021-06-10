package leetcode

import (
	"math"
	"sort"
)

// solution1
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result, difference := 0, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				sum = nums[i] + nums[j] + nums[k]
				if abs(sum-target) < difference {
					difference = abs(sum - target)
					result = sum
				}
			}
		}
	}
	return result
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// solution2
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result, difference := 0, math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		sum := 0
		for j, k := i+1, len(nums)-1; j < k; {
			sum = nums[i] + nums[j] + nums[k]
			if abs(sum-target) < difference {
				difference = abs(sum - target)
				result = sum
			}
			if sum == target {
				return sum
			} else if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return result
}
