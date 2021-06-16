package leetcode

// solution1
func swapNumbers(numbers []int) []int {
	numbers[0], numbers[1] = numbers[1], numbers[0]
	return numbers
}

// solution1
func swapNumbers(numbers []int) []int {
	numbers[0] ^= numbers[1]
	numbers[1] ^= numbers[0]
	numbers[0] ^= numbers[1]
	return numbers
}
