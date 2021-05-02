package leetcode

func leastBricks(wall [][]int) int {
	widthMap := make(map[int]int)
	for _, widths := range wall {
		sum := 0
		for width := 0; width < len(widths)-1; width++ {
			sum += widths[width]
			widthMap[sum]++
		}
	}
	maxWidth := 0
	for _, v := range widthMap {
		maxWidth = max(maxWidth, v)
	}
	return len(wall) - maxWidth
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
