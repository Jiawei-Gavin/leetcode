package leetcode

import "strconv"

func maxPoints(points [][]int) int {
	if len(points) < 3 {
		return len(points)
	}
	var res int
	for i := 0; i < len(points); i++ {
		tmp, cnt, hashMap := 0, 0, make(map[string]int)
		for j := i + 1; j < len(points); j++ {
			x, y := points[j][0]-points[i][0], points[j][1]-points[i][1]
			if x == 0 && y == 0 {
				cnt++
				continue
			}
			divisor := gcd(x, y)
			x, y = x/divisor, y/divisor
			hashMap[strconv.Itoa(x)+"/"+strconv.Itoa(y)]++
			tmp = max(tmp, hashMap[strconv.Itoa(x)+"/"+strconv.Itoa(y)])
		}
		res = max(res, tmp+cnt+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
