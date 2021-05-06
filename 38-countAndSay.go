package leetcode

func countAndSay(n int) string {
	result := []byte{'1'}
	for i := 1; i < n; i++ {
		var temp []byte
		for j := 0; j < len(result); {
			count := 1
			for k := j + 1; k < len(result); k++ {
				if result[j] == result[k] {
					count++
				} else {
					break
				}
			}
			temp = append(temp, byte(count+'0'), result[j])
			j += count
		}
		result = temp
	}
	return string(result)
}
