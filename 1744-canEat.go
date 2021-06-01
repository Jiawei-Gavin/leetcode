package leetcode

func canEat(candiesCount []int, queries [][]int) []bool {
	result := make([]bool, len(queries))
	sum := make(map[int]int)
	sum[0] = candiesCount[0]
	for i := 1; i < len(candiesCount); i++ {
		sum[i] = sum[i-1] + candiesCount[i]
	}

	for i, q := range queries {
		favoriteType, favoriteDay, dailyCap := q[0], q[1], q[2]
		result[i] = favoriteDay+1 <= sum[favoriteType] && (favoriteDay+1)*dailyCap > sum[favoriteType-1]
	}
	return result
}
