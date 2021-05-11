package leetcode

func subsets(nums []int) [][]int {
	var temp []int
	var result [][]int
	subsetsFunc(temp, nums, &result, 0)
	return result
}

func subsetsFunc(temp []int, nums []int, result *[][]int, start int) {
	track := make([]int, len(temp))
	copy(track, temp)
	*result = append(*result, track)

	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		subsetsFunc(temp, nums, result, i+1)
		temp = temp[:len(temp)-1]
	}
}
