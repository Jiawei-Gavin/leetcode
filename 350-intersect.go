package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	hashTable := make(map[int]int)
	var result []int
	for _, num := range nums1 {
		hashTable[num]++
	}
	for _, num := range nums2 {
		if hashTable[num] > 0 {
			result = append(result, num)
			hashTable[num]--
		}
	}
	return result
}
