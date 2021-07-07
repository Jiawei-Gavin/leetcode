package leetcode

import (
	"strconv"
	"strings"
)

// solution1
func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i := 0; i < max(len(v1), len(v2)); i++ {
		temp1, temp2 := 0, 0
		if i < len(v1) {
			temp1, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			temp2, _ = strconv.Atoi(v2[i])
		}
		if temp1 == temp2 {
			continue
		}
		if temp1 < temp2 {
			return -1
		}
		return 1
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// solution2
func compareVersion(version1 string, version2 string) int {
	m, n := len(version1), len(version2)
	i, j := 0, 0
	for i < m || j < n {
		v1, v2 := 0, 0
		for i < m && version1[i] != '.' {
			v1 = v1*10 + int(version1[i]-'0')
			i++
		}
		for j < n && version2[j] != '.' {
			v2 = v2*10 + int(version2[j]-'0')
			j++
		}
		if v1 != v2 {
			if v1 > v2 {
				return 1
			}
			return -1
		}
		i++
		j++
	}
	return 0
}
