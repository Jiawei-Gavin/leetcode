package leetcode

func countTriplets(arr []int) int {
	cnt, total, s, result := map[int]int{}, map[int]int{}, 0, 0
	for k, val := range arr {
		if i, ok := cnt[s^val]; ok {
			result += i*k - total[s^val]
		}
		cnt[s]++
		total[s] += k
		s ^= val
	}
	return result
}
