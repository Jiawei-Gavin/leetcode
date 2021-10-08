package leetcode

func findRepeatedDnaSequences(s string) []string {
	var res []string
	visited := make(map[string]int)
	for i, j := 0, 10; j <= len(s); i, j = i+1, j+1 {
		if visited[s[i:j]] == 1 {
			res = append(res, s[i:j])
		}
		visited[s[i:j]]++
	}
	return res
}
