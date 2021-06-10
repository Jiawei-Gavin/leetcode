package leetcode

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	result := make([]int, len(this.nums))
	copy(result, this.nums)
	for i := len(this.nums) - 1; i >= 0; i-- {
		randIdx := rand.Intn(i + 1)
		result[i], result[randIdx] = result[randIdx], result[i]
	}
	return result
}
