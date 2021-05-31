package leetcode

// solution1
func singleNumber(nums []int) int {
	hashTable := make(map[int]int)
	for _, num := range nums {
		hashTable[num]++
	}
	for num, val := range hashTable {
		if val == 1 {
			return num
		}
	}
	return 0
}

// solution2
func singleNumber(nums []int) int {
	var res int
	for i := 0; i < 32; i++ {
		tmp := 0
		for j := 0; j < len(nums); j++ {
			tmp += (nums[j] >> i) & 1
		}
		res |= (tmp % 3) << i
	}
	return res
}

// solution3
func singleNumber(nums []int) int {
	m, n := 0, 0
	for _, num := range nums {
		m = m ^ num & ^n
		n = n ^ num & ^m
	}
	return m
}
