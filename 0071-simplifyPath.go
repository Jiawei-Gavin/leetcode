package leetcode

import "strings"

func simplifyPath(path string) string {
	str := strings.Split(path, "/")
	var stack []string
	for _, s := range str {
		switch s {
		case "", ".":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, s)
		}
	}
	return "/" + strings.Join(stack, "/")
}
