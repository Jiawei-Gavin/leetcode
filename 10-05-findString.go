package leetcode

func findString(words []string, s string) int {
	left, right := 0, len(words)-1
	for left <= right {
		mid := (left + right) >> 1
		for ; mid < right; mid++ {
			if words[mid] != "" {
				break
			}
		}
		if words[mid] == s {
			return mid
		}
		if mid == right {
			right = (left+right)>>1 - 1
			continue
		}
		if words[mid] > s {
			right = mid - 1
		} else if words[mid] < s {
			left = mid + 1
		}
	}
	return -1
}
