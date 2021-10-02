package leetcode

func destCity(paths [][]string) string {
	hashMap := make(map[string]struct{})
	for _, path := range paths {
		hashMap[path[0]] = struct{}{}
	}
	for _, path := range paths {
		if _, has := hashMap[path[1]]; !has {
			return path[1]
		}
	}
	return ""
}
