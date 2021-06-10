package leetcode

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	var result []byte
	if numerator*denominator < 0 {
		result = append(result, '-')
	}

	numerator = abs(numerator)
	denominator = abs(denominator)
	result = append(result, []byte(strconv.Itoa(numerator/denominator))...)
	numerator = numerator % denominator
	if numerator == 0 {
		return string(result)
	}

	result = append(result, '.')
	var mem = make(map[int]int)
	var start = -1
	var lowRes []byte
	for numerator > 0 {
		if pos, ok := mem[numerator]; ok {
			start = pos
			break
		}
		mem[numerator] = len(lowRes)
		numerator = numerator * 10
		lowRes = append(lowRes, []byte(strconv.Itoa(numerator/denominator))...)
		numerator = numerator % denominator
	}
	if start > -1 {
		result = append(result, lowRes[:start]...)
		result = append(result, '(')
		result = append(result, lowRes[start:]...)
		result = append(result, ')')
	} else {
		result = append(result, lowRes...)
	}

	return string(result)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
