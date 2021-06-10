package leetcode

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		result := 0
		if height[end] < height[start] {
			result = height[end] * width
			end--
		} else {
			result = height[start] * width
			start++
		}
		if result > max {
			max = result
		}
	}
	return max
}
