package leetcode

// solution1
func merge(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n; p > 0; p-- {
		if m > 0 && n > 0 {
			if nums1[m-1] <= nums2[n-1] {
				nums1[p-1] = nums2[n-1]
				n--
			} else {
				nums1[p-1] = nums1[m-1]
				m--
			}
		}
	}

	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}

// solution2
func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, m+n)
	p1, p2 := 0, 0
	for p := 0; p < m+n; p++ {
		if p1 == m {
			for p1+p2 < m+n {
				sorted[p] = nums2[p2]
				p2++
				p++
			}
			break
		}
		if p2 == n {
			for p1+p2 < m+n {
				sorted[p] = nums1[p1]
				p1++
				p++
			}
			break
		}
		if nums1[p1] > nums2[p2] {
			sorted[p] = nums2[p2]
			p2++
		} else {
			sorted[p] = nums1[p1]
			p1++
		}
	}
	copy(nums1, sorted)
}
