package leetcode

var roman = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	n := roman[s[len(s)-1]]
	for i := len(s) - 1; i > 0; i-- {
		preNum := roman[s[i-1]]
		if roman[s[i]] > preNum {
			n -= preNum
		} else {
			n += preNum
		}
	}
	return n
}
