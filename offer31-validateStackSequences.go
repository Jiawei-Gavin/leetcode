package leetcode

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for popped[0] == stack[len(stack)-1] {
			popped = popped[1:]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
		}
	}
	return len(popped) == 0
}
