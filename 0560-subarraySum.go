package leetcode

func subarraySum(nums []int, k int) int {
	result, sum, numMap := 0, 0, map[int]int{}
	numMap[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := numMap[sum-k]; ok {
			result += numMap[sum-k]
		}
		numMap[sum] += 1
	}
	return result
}
