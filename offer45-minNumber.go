package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	var str []string
	for _, num := range nums {
		str = append(str, strconv.Itoa(num))
	}
	sort.Slice(str, func(i, j int) bool {
		if str[i]+str[j] < str[j]+str[i] {
			return true
		}
		return false
	})
	return strings.Join(str, "")
}
