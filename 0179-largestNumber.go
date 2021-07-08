package leetcode

import (
	"fmt"
	"sort"
	"strings"
)

func largestNumber(nums []int) string {
	str := make([]string, len(nums))
	for i, _ := range nums {
		str[i] = fmt.Sprintf("%d", nums[i])
	}
	sort.Slice(str, func(i, j int) bool {
		return str[i]+str[j] > str[j]+str[i]
	})
	if str[0][0] == '0' {
		return "0"
	}
	return strings.Join(str, "")
}
