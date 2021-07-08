package leetcode

func findRepeatedDnaSequences(s string) []string {
	var res []string
	visited := make(map[string]int)
	for i := 0; i < len(s)-10+1; i++ {
		str := s[i : i+10]
		if _, ok := visited[str]; ok {
			res = append(res, str)
		}
		visited[str]++
	}
	return res
}
