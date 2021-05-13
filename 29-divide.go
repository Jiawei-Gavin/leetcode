package leetcode

import "math"

// solution1
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	result := 0
	sign := -1
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		sign = 1
	}
	dvd, dvs := abs(dividend), abs(divisor)
	for dvd >= dvs {
		temp := dvs
		m := 1
		for temp<<1 <= dvd {
			temp <<= 1
			m <<= 1
		}
		dvd -= temp
		result += m
	}
	return sign * result
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// solution2
func divide(dividend int, divisor int) int {
	sign := true
	if dividend < 0 {
		dividend = 0 - dividend
		sign = !sign
	}
	if divisor < 0 {
		divisor = 0 - divisor
		sign = !sign
	}
	if divisor > dividend {
		return 0
	}
	result, mulDivisor := 1, divisor
	for dividend > mulDivisor+mulDivisor {
		result += result
		mulDivisor += mulDivisor
	}
	result += divide(dividend-mulDivisor, divisor)
	if !sign {
		result = 0 - result
	}
	if result > 1<<31-1 || 0-result > 1<<31 {
		return 1<<31 - 1
	}
	return result
}
