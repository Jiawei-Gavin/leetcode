package leetcode

func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m*k {
		return -1
	}

	max := 0
	for i := range bloomDay {
		if bloomDay[i] > max {
			max = bloomDay[i]
		}
	}

	min := 0
	for min < max {
		mid := (min + max) / 2
		flowers, bouquets := 0, 0
		for i := 0; i < len(bloomDay); i++ {
			flowers++
			if bloomDay[i] > mid {
				flowers = 0
			}
			if flowers == k {
				bouquets++
				flowers = 0
			}
		}
		if bouquets >= m {
			max = mid
		} else {
			min = mid + 1
		}
	}
	return min
}
