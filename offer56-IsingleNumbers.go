package leetcode

func singleNumbers(nums []int) []int {
	num1, num2, m, n := 0, 0, 1, 0
	for _, num := range nums {
		n ^= num
	}
	m = n & (-n)
	for _, num := range nums {
		if num&m != 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}
	return []int{num1, num2}
}
