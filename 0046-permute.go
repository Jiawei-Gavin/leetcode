package leetcode

func permute(nums []int) [][]int {
	var pathNums []int
	var used = make([]bool, len(nums))
	var result [][]int
	permutationFunc(nums, pathNums, &result, &used)
	return result
}

func permutationFunc(nums, pathNums []int, result *[][]int, used *[]bool) {
	if len(nums) == len(pathNums) {
		tmp := make([]int, len(nums))
		copy(tmp, pathNums)
		*result = append(*result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			pathNums = append(pathNums, nums[i])
			permutationFunc(nums, pathNums, result, used)
			pathNums = pathNums[:len(pathNums)-1]
			(*used)[i] = false
		}
	}
}
