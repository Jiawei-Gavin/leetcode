package leetcode

import (
	"math"
	"strconv"
)

func nextGreaterElement(n int) int {
	arr := []byte(strconv.Itoa(n))
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] < arr[i+1] {
			for j := len(arr) - 1; j > i; j-- {
				if arr[j] > arr[i] {
					arr[i], arr[j] = arr[j], arr[i]
					break
				}
			}
			for j := i + 1; j <= (len(arr)-i)/2+i; j++ {
				arr[j], arr[len(arr)-j+i] = arr[len(arr)-j+i], arr[j]
			}
			res, _ := strconv.Atoi(string(arr[:]))
			if res > math.MaxInt32 {
				return -1
			}
			return res
		}
	}
	return -1
}
