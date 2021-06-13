package leetcode

func printBin(num float64) string {
	res := "0."
	for num != 0 {
		num *= 2
		if num >= 1 {
			res += "1"
			num -= 1
		} else {
			res += "0"
		}
	}
	if len(res) >= 32 {
		return "ERROR"
	}
	return res
}
