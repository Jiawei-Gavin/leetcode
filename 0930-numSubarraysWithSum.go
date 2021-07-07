package leetcode

func numSubarraysWithSum(nums []int, goal int) int {
	hashMap := make(map[int]int)
	var res int
	sum := 0
	for _, num := range nums {
		hashMap[sum]++
		sum += num
		res += hashMap[sum-goal]
	}
	return res
}
