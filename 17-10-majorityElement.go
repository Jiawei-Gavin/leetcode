package leetcode

func majorityElement(nums []int) int {
	prev, count := -1, 0
	for _, num := range nums {
		if count == 0 {
			count++
			prev = num
		} else if prev == num {
			count++
		} else {
			count--
		}
	}
	if count > 0 {
		n := 0
		for _, num := range nums {
			if num == prev {
				n++
			}
			if n > len(nums)/2 {
				return prev
			}
		}
	}
	return -1
}
