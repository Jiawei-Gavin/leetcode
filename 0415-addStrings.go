package leetcode

import "strconv"

func addStrings(num1 string, num2 string) string {
	var res string
	carry := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
		var temp1, temp2 int
		if i >= 0 {
			temp1 = int(num1[i] - '0')
		}
		if j >= 0 {
			temp2 = int(num2[j] - '0')
		}
		temp := temp1 + temp2 + carry
		res = strconv.Itoa(temp%10) + res
		carry = temp / 10
	}
	return res
}
