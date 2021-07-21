package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length%2 == 1 {
		mid := length / 2
		return getKthElement(nums1, nums2, mid+1)
	} else {
		mid1, mid2 := length/2-1, length/2
		return (getKthElement(nums1, nums2, mid1+1) + getKthElement(nums1, nums2, mid2+1)) / 2
	}
}

func getKthElement(nums1 []int, nums2 []int, k int) float64 {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return float64(nums2[index2+k-1])
		}
		if index2 == len(nums2) {
			return float64(nums1[index1+k-1])
		}
		if k == 1 {
			return float64(min(nums1[index1], nums2[index2]))
		}
		mid := k / 2
		newIndex1 := min(index1+mid, len(nums1)) - 1
		newIndex2 := min(index2+mid, len(nums2)) - 1
		if nums1[newIndex1] <= nums2[newIndex2] {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
