package leetcode

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	hashTable := make(map[int]int)
	var res int
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			hashTable[num1+num2]++
		}
	}
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			res += hashTable[-num3-num4]
		}
	}
	return res
}
