package leetcode

import "sort"

func permutation(S string) []string {
	if len(S) == 0 {
		return nil
	}
	var str []string
	for _, value := range S {
		str = append(str, string(value))
	}
	var result []string
	sort.Strings(str)
	dfs(str, 0, &result)
	return result
}

func dfs(str []string, i int, result *[]string) {
	n := len(str)
	if i == n-1 {
		var tmp string
		for _, s := range str {
			tmp += s
		}
		*result = append(*result, tmp)
		return
	}
	for j := i; j < n; j++ {
		if j != i && str[j] == str[i] {
			continue
		}
		str[j], str[i] = str[i], str[j]
		dfs(str, i+1, result)
	}
	for k := n - 1; k > i; k-- {
		str[i], str[k] = str[k], str[i]
	}
}
