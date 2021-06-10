package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]int)
	var result []int
	for _, num := range nums1 {
		hashMap[num]++
	}
	for _, num := range nums2 {
		if hashMap[num] > 0 {
			result = append(result, num)
			hashMap[num]--
		}
	}
	return result
}
