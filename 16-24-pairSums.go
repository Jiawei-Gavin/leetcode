package leetcode

func pairSums(nums []int, target int) [][]int {
	var res [][]int
	hashMap := make(map[int]int)
	for _, num := range nums {
		if cnt, ok := hashMap[target-num]; ok && cnt > 0 {
			res = append(res, []int{target - num, num})
			hashMap[target-num]--
		} else {
			hashMap[num]++
		}
	}
	return res
}
