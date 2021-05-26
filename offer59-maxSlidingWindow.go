package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	window := make([]int, 0, k)
	result := make([]int, 0, len(nums)-k+1)
	for i, num := range nums {
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < num {
			window = window[0 : len(window)-1]
		}
		window = append(window, i)
		if i >= k-1 {
			result = append(result, nums[window[0]])
		}
	}
	return result
}
