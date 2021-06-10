package leetcode

func permutation(s string) []string {
	var result []string
	bytes := []byte(s)
	dfs(0, bytes, &result)
	return result
}

func dfs(n int, bytes []byte, result *[]string) {
	if n == len(bytes)-1 {
		*result = append(*result, string(bytes))
	}
	hashMap := map[byte]bool{}
	for i := n; i < len(bytes); i++ {
		if !hashMap[bytes[i]] {
			bytes[n], bytes[i] = bytes[i], bytes[n]
			hashMap[bytes[n]] = true
			dfs(n+1, bytes, result)
			bytes[n], bytes[i] = bytes[i], bytes[n]
		}
	}
}
