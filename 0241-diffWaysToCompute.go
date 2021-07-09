package leetcode

import "strconv"

func diffWaysToCompute(expression string) []int {
	var res []int
	if len(expression) == 0 {
		return res
	}
	for i := 0; i < len(expression); i++ {
		temp := expression[i]
		if temp == '+' || temp == '-' || temp == '*' {
			left := diffWaysToCompute(expression[0:i])
			right := diffWaysToCompute(expression[i+1 : len(expression)])
			for _, leftNum := range left {
				for _, rightNum := range right {
					if temp == '-' {
						res = append(res, leftNum-rightNum)
					} else if temp == '+' {
						res = append(res, leftNum+rightNum)
					} else {
						res = append(res, leftNum*rightNum)
					}
				}
			}
		}
	}
	if len(res) == 0 {
		temp, _ := strconv.Atoi(expression)
		res = append(res, temp)
	}
	return res
}
