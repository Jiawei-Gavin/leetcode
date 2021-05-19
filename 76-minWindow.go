package leetcode

import "math"

func minWindow(s string, t string) string {
	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	match := 0
	index, length := 0, math.MaxInt32
	for right < len(s) {
		tempAdd := s[right]
		right++
		if _, ok := need[tempAdd]; ok {
			window[tempAdd]++
			if window[tempAdd] == need[tempAdd] {
				match++
			}
		}
		for match == len(need) {
			if right-left < length {
				index = left
				length = right - left
			}
			tempDel := s[left]
			left++

			if _, ok := need[tempDel]; ok {
				window[tempDel]--
				if window[tempDel] < need[tempDel] {
					match--
				}
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[index : index+length]
}
