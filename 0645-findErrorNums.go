package leetcode

func findErrorNums(nums []int) []int {
	hashMap := make(map[int]int)
	for _, v := range nums {
		hashMap[v]++
	}
	res := make([]int, 2)
	for i := 1; i <= len(nums); i++ {
		if hashMap[i] == 2 {
			res[0] = i
		} else if hashMap[i] == 0 {
			res[1] = i
		}
	}
	return res
}
