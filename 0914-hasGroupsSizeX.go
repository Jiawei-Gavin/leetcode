package leetcode

func hasGroupsSizeX(deck []int) bool {
	cnt := make(map[int]int)
	for _, num := range deck {
		cnt[num]++
	}
	gcdNum := cnt[deck[0]]
	for _, num := range cnt {
		gcdNum = gcd(gcdNum, num)
		if gcdNum < 2 {
			return false
		}
	}
	return gcdNum >= 2
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
