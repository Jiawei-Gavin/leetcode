package leetcode

func longestConsecutive(nums []int) int {
	result, numMap := 0, map[int]bool{}
	for _, num := range nums {
		numMap[num] = true
	}
	for num := range numMap {
		if !numMap[num-1] {
			curNum := num
			sum := 1
			for numMap[curNum+1] {
				curNum++
				sum++
			}
			if result < sum {
				result = sum
			}
		}
	}
	return result
}
