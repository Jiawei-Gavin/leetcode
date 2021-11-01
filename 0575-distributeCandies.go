package leetcode

func distributeCandies(candyType []int) int {
	set := map[int]struct{}{}
	for _, candy := range candyType {
		set[candy] = struct{}{}
	}
	if len(candyType)/2 <= len(hashMap) {
		return len(candyType) / 2
	}
	return len(hashMap)
}
