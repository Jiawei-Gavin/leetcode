package leetcode

func decodeString(s string) string {
	var stackStr []string
	var stackNum []int
	num := 0
	var temp string
	for i, v := range s {
		if v >= '0' && v <= '9' {
			num = num*10 + int(v-'0')
		} else if v == '[' {
			stackNum = append(stackNum, num)
			stackStr = append(stackStr, temp)
			num = 0
			temp = ""
		} else if v == ']' {
			count := stackNum[len(stackNum)-1]
			stackNum = stackNum[:len(stackNum)-1]
			for i := 0; i < count; i++ {
				stackStr[len(stackStr)-1] += temp
			}
			temp = stackStr[len(stackStr)-1]
			stackStr = stackStr[:len(stackStr)-1]
		} else {
			temp += string(s[i])
		}
	}
	if len(stackStr) == 0 {
		return temp
	}
	return stackStr[len(stackStr)-1]
}
