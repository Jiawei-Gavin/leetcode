package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	countZero, flower := 1, 0
	for _, bed := range flowerbed {
		if bed == 0 {
			countZero++
		} else {
			flower += (countZero - 1) / 2
			if flower >= n {
				return true
			}
			countZero = 0
		}
	}
	countZero++
	flower += (countZero - 1) / 2
	return flower >= n
}
