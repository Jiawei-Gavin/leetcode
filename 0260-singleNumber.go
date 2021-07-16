package leetcode

func singleNumber(nums []int) []int {
	res := make([]int, 2)
	rest := 0
	for i := 0; i < len(nums); i++ {
		rest ^= nums[i]
	}
	rest &= -rest
	for _, num := range nums {
		if num&rest == 0 {
			res[0] ^= num
		} else {
			res[1] ^= num
		}
	}
	return res
}
