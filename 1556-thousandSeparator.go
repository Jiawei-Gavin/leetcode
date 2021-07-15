package leetcode

import "strconv"

func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}
	count := 0
	var res string
	for n != 0 {
		count++
		temp := n % 10
		n /= 10
		res += strconv.Itoa(temp)
		if count == 3 && n != 0 {
			res += "."
			count = 0
		}
	}
	return reverse(res)
}

func reverse(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
