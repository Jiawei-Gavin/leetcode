package leetcode

func findSwapValues(array1 []int, array2 []int) []int {
	sum1, sum2 := 0, 0
	for _, array := range array1 {
		sum1 += array
	}
	hashMap := make(map[int]struct{})
	for _, array := range array2 {
		sum2 += array
		hashMap[array] = struct{}{}
	}
	diff := sum1 - sum2
	if diff%2 == 0 {
		diff >>= 1
	} else {
		return []int{}
	}
	for _, v1 := range array1 {
		tmp := v1 - diff
		if _, ok := hashMap[tmp]; ok {
			return []int{v1, tmp}
		}
	}
	return []int{}
}
