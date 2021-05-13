package leetcode

// solution1
func isHappy(n int) bool {
	sum, num, record := 0, n, map[int]int{}
	for {
		for num != 0 {
			sum += (num % 10) * (num % 10)
			num = num / 10
		}
		if _, ok := record[sum]; !ok {
			if sum == 1 {
				return true
			}
			record[sum] = sum
			num = sum
			sum = 0
			continue
		} else {
			return false
		}
	}
}

// solution2
func isHappy(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && fast != slow {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	result := 0
	for n > 0 {
		result += (n % 10) * (n % 10)
		n /= 10
	}
	return result
}
