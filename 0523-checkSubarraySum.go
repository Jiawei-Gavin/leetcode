package leetcode

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	hashMap := map[int]int{0: -1}
	remainder := 0
	for i, num := range nums {
		remainder = (remainder + num) % k
		if index, ok := hashMap[remainder]; ok {
			if i-index >= 2 {
				return true
			}
		} else {
			hashMap[remainder] = i
		}
	}
	return false
}
