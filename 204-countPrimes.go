package leetcode

func countPrimes(n int) int {
	isNotPrime := make([]bool, n)
	for i := range isNotPrime {
		isNotPrime[i] = true
	}
	count := 0
	for i := 2; i < n; i++ {
		if isNotPrime[i] {
			count++
			for j := 2 * i; j < n; j += i {
				isNotPrime[j] = false
			}
		}
	}
	return count
}
