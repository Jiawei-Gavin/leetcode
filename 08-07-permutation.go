package leetcode

func permutation(S string) []string {
	var res []string
	bytes := []byte(S)
	dfs(0, bytes, &res)
	return res
}

func dfs(n int, bytes []byte, res *[]string) {
	if n == len(bytes)-1 {
		*res = append(*res, string(bytes))
	}
	hashMap := map[byte]bool{}
	for i := n; i < len(bytes); i++ {
		if !hashMap[bytes[i]] {
			bytes[n], bytes[i] = bytes[i], bytes[n]
			hashMap[bytes[n]] = true
			dfs(n+1, bytes, res)
			bytes[n], bytes[i] = bytes[i], bytes[n]
		}
	}
}
