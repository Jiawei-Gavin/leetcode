package leetcode

func maxArea(height []int) int {
	res, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		area := 0
		if height[end] < height[start] {
			area = height[end] * width
			end--
		} else {
			area = height[start] * width
			start++
		}
		res = max(res, area)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
